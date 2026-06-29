package handler

import (
	"errors"
	"net/http"

	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC usecase.AuthUseCase
}

func NewAuthHandler(authUC usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		authUC: authUC,
	}
}
func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest
}