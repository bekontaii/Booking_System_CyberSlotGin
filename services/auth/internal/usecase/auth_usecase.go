package usecase

import (
	"context"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/repository"
)

type AuthUseCase interface {
	Register(ctx context.Context, user *domain.RegisterRequest) error
	Login(ctx context.Context, user *domain.LoginRequest) (string, error)
}
type authUseCase struct {
	repo repository.UserRepository
}
