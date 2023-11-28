package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guvanchhojamov/app-todo/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAllItems)
				items.POST("/", h.createItem)
			}

		}
		items := api.Group("/items")
		{
			items.GET("/:item_id", h.getItemById)
			items.PUT("/:item_id", h.updateItem)
			items.DELETE("/:item_id", h.deleteItem)
		}
	}
	return router
}
