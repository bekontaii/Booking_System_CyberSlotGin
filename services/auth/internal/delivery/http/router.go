package http

import (
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/delivery/http/handlers"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(router *gin.RouterGroup, authUC usecase.AuthUseCase) {
	authHandler := handlers.NewAuthHandler(authUC)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
}
