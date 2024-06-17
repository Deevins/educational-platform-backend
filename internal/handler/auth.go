package handler

import (
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.UserCreate

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if input.Email == "" || input.Password == "" {
		model.NewErrorResponse(ctx, http.StatusBadRequest, "email or password is empty")
		return
	}

	res, err := h.as.CreateUser(ctx, input)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":    res.UserID,
		"role":  "USER",
		"token": res.Token,
	})

}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(ctx *gin.Context) {
	var input signInInput

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.as.Authorize(ctx, input.Email, input.Password)
	if err != nil {
		model.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":    res.UserID,
		"token": res.Token,
	})
}
