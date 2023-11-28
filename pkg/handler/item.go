package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"net/http"
	"strconv"
)

type GetAllItemsData struct {
	UserId int
	listId int
	Data   []model.TodoItem
}

func (h *Handler) createItem(ctx *gin.Context) {
	_, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount CreateItem")
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param: "+err.Error())
		return
	}

	var input model.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.CreateItem(input, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "todoItem.CreateItem new item Id error")
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"item id": itemId,
	})
}

func (h *Handler) getAllItems(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount CreateItem")
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	fmt.Printf("List Id: %d \n", listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param: "+err.Error())
		return
	}
	items, err := h.services.TodoItem.GetUserAllItems(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "handler GetUserAllItems error "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, GetAllItemsData{
		UserId: userId,
		listId: 1,
		Data:   items,
	})
}

func (h *Handler) getItemById(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount CreateItem")
		return
	}
	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	fmt.Printf("Item Id: %d \n", itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param item id: "+err.Error())
		return
	}

	itemById, err := h.services.TodoItem.GetItemById(userId, itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "handler GetItemById error : "+err.Error())
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"itemById": itemById,
		"itemId":   itemId,
		"userId":   userId,
	})

}

func (h *Handler) updateItem(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not found Update Item")
		return
	}

	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	fmt.Printf("List Id: %d \n", itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param: "+err.Error())
		return
	}

	var input model.UpdateItemInput
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "handler Bind input error UpdateItem: "+err.Error())
		return
	}

	if err = h.services.TodoItem.UpdateItem(input, userId, itemId); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "handler Update Item error "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"User Id": userId,
		"Item Id": itemId,
		"updated": "ok",
	})
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount CreateItem")
		return
	}
	itemId, err := strconv.Atoi(ctx.Param("item_id"))
	fmt.Printf("Item Id: %d \n", itemId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param item id: "+err.Error())
		return
	}
	if err := h.services.TodoItem.DeleteItem(userId, itemId); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "handler delete item error: "+err.Error())
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"User Id": userId,
		"itemId":  itemId,
		"deleted": "ok",
	})
}
