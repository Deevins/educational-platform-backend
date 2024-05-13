package handler

import (
	"context"
	"fmt"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthService interface {
	CreateUser(ctx context.Context, user model.UserCreate) (int32, error)
}

type Handler struct {
	as AuthService
}

func NewHandler(authSvc AuthService) *Handler {
	return &Handler{as: authSvc}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		//auth.POST("/sign-in", h.signIn)
	}

	//api := router.Group("/api", h.userIdentity)
	//{
	//
	//
	//	items := api.Group("items")
	//	{
	//		items.GET("/:id", h.getItemById)
	//		items.PUT("/:id", h.updateItem)
	//		items.DELETE("/:id", h.deleteItem)
	//	}
	//}
	return router
}

func (h *Handler) signUp(ctx *gin.Context) {
	var input model.UserCreate

	if err := ctx.BindJSON(&input); err != nil {
		fmt.Println(err)
		//NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.as.CreateUser(ctx, input)
	if err != nil {
		fmt.Println(err)
		//NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
