package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"net/http"
	"strconv"
)

type GetAllListsData struct {
	Data []model.TodoList `json:"data"`
}

func (h *Handler) getAllLists(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount GetAllLists")
		return
	}

	lists, err := h.services.TodoList.GetAllList(userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "todoList.GetAlllist listId error: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, GetAllListsData{
		Data: lists,
	})

}

func (h *Handler) createList(ctx *gin.Context) {

	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount CreateList")
		return
	}

	var input model.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "todoList.Create listId error")
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"listId": listId,
	})

}

func (h *Handler) getListById(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount GetListById: "+err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param: "+err.Error())
		return
	}
	listData, err := h.services.TodoList.GetListById(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "todoList.GetListById listData error: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, listData)

}

func (h *Handler) deleteList(ctx *gin.Context) {
	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount GetListById: "+err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param: "+err.Error())
		return
	}

	//call Delete services
	err = h.services.TodoList.DeleteList(userId, listId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "todoList Delete list error: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		UserId: userId,
		Status: "Ok",
	})

}

func (h *Handler) updateList(ctx *gin.Context) {

	userId, err := GetUserId(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, "id not fount GetListById: "+err.Error())
		return
	}

	listId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid id Param: "+err.Error())
		return
	}

	var input model.UpdateListInput
	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "update list bind fileds error: "+err.Error())
		return
	}

	err = h.services.TodoList.UpdateList(input, listId, userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "update list UpdateList error: "+err.Error())
		return
	}

	ctx.JSON(http.StatusOK, statusResponse{
		Status: "ok",
		UserId: userId,
	})
}
