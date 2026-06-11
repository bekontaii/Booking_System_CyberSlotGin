package main

import (
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/config"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/delivery/http"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/repository"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()

	userRepo := repository.NewInMemoryUserRepository()

	authUC := usecase.NewAuthUseCase(userRepo, cfg.JWT.Secret, cfg.JWT.TTL)

	router := gin.Default()
	http.MapAuthRoutes(router.Group("/api"), authUC)

	router.Run(":" + cfg.Server.Port)
}