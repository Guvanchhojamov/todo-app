package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	headerToken         = "Token"
	userCtx             = "UserId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(headerToken)
	if header == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "User Not Authorized! Header invalid")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(ctx, http.StatusUnauthorized, "User Not Authorized! Header Auth is invalid")
		return
	}
	fmt.Println(header) // Proverka headera
	UserId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}
	ctx.Set(userCtx, UserId)
}

func GetUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "user id not found GetUserId")
		return 0, errors.New("user id not found GetUserID")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(ctx, http.StatusInternalServerError, "error convert user id to int GetUserId")
		return 0, errors.New("error convert user id to int GetUserId")
	}
	return idInt, nil
}
