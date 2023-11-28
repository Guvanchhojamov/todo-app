package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guvanchhojamov/app-todo/pkg/model"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
