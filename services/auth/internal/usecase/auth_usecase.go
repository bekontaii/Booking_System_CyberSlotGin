package usecase

import (
	"context"
	"errors"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (a *authUseCase) Register(ctx context.Context, user *domain.RegisterRequest) error {
	_, err := a.repo.GetUserByUsername(ctx, user.Username)
	if err == nil {
		return ErrUsernameAlreadyExists
	}

	if !errors.Is(err, repository.ErrUserNotFound) {
		return err
	}
	_, err = a.repo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return ErrEmailAlreadyExists
	}
	if !errors.Is(err, repository.ErrUserNotFound) {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	createdUser := &domain.User{
		Username:     user.Username,
		Name:         user.Name,
		Surname:      user.Surname,
		Email:        user.Email,
		PasswordHash: string(hashedPassword),
		Role:         "User",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = a.repo.CreateUser(ctx, createdUser)
	if err != nil {
		return err
	}
	return nil
}
func (a *authUseCase) Login(ctx context.Context, user *domain.LoginRequest) (string, error) {

}
