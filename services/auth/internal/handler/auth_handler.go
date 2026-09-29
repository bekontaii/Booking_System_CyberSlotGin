package handler

import (
	"errors"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	clubUC usecase.AuthUseCase
}

func NewAuthHandler(clubUC usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		clubUC: clubUC,
	}
}
func (h *AuthHandler) Register(c *gin.Context) {
	var req domain.RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status Bad Request"})
		return
	}
	err = h.authUC.Register(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, usecase.ErrUsernameAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
		if errors.Is(err, usecase.ErrEmailAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}
func (h *AuthHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status Bad Request"})
		return
	}
	token, err := h.authUC.Login(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, "token:"+token)
}
