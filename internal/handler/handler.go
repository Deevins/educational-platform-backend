package handler

import (
	"context"
	"github.com/deevins/educational-platform-backend/internal/model"
	"github.com/deevins/educational-platform-backend/internal/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthService interface {
	CreateUser(ctx context.Context, user model.UserCreate) (user_service.RegisterUserResponse, error)
	GenerateToken(ctx context.Context, email, password string) (user_service.LoginUserResponse, error)
}

type UserService interface {
	GetByID(ctx context.Context, ID int32) (*model.User, error)
}

type Handler struct {
	as AuthService

	us UserService
}

func NewHandler(authSvc AuthService) *Handler {
	return &Handler{as: authSvc}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(corsMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	user := router.Group("/user")
	{
		user.GET("/get-one", h.getOneUser)
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

func (h *Handler) getOneUser(ctx *gin.Context) {
	var input int32 // id

	if err := ctx.BindJSON(&input); err != nil {
		model.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
