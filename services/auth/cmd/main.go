package main

import (
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/handler"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/repository"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/routes"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	repo := repository.NewMemoryUserRepository()
	secret := []byte(os.Getenv("JWT_SECRET"))
	uc := usecase.NewAuthUseCase(repo, secret)
	handler := handler.NewAuthHandler(&uc)
	routes.RegisterRoutes(r, handler)
	r.Run(":8080")
}
