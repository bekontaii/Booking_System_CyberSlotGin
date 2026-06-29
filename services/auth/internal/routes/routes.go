package routes

import (
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
}
