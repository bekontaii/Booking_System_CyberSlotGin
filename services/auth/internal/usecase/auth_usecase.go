package usecase

import (
	"context"
	"errors"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/repository"
)

type AuthUseCase interface {
	Register(ctx context.Context, user *domain.RegisterRequest) error
	Login(ctx context.Context, user *domain.LoginRequest) (string, error)
}
type authUseCase struct {
	repo      repository.UserRepository
	jwtSecret []byte
}

func NewAuthUseCase(repo repository.UserRepository, jwtSecret []byte) *authUseCase {
	return &authUseCase{repo: repo, jwtSecret: jwtSecret}
}
func (a authUseCase) Register(ctx context.Context, user *domain.RegisterRequest) (string, error) {
	if user.Username == "" || user.Password == "" {
		return " ", errors.New("username or password is empty")
	}
	if user.Email == "" || user.Password == "" {
	}
}
