package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4" 
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
)

type AuthUseCase interface {
	Register(ctx context.Context, req *domain.RegisterRequest) error
	Login(ctx context.Context, req *domain.LoginRequest) (string, error)
}

type authUseCase struct {
	userRepo  repository.UserRepository
	jwtSecret []byte
	tokenTTL  time.Duration 
}


func NewAuthUseCase(repo repository.UserRepository, secret string, ttl time.Duration) AuthUseCase {
	return &authUseCase{
		userRepo:  repo,
		jwtSecret: []byte(secret),
		tokenTTL:  ttl,
	}
}


func (u *authUseCase) Register(ctx context.Context, req *domain.RegisterRequest) error {
	existingUser, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err == nil && existingUser != nil {
		return repository.ErrUsernameExists
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}


	newUser := &domain.User{
		Name:         req.Name,
		Surname:      req.Surname,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}


	return u.userRepo.CreateUser(ctx, newUser)
}


func (u *authUseCase) Login(ctx context.Context, req *domain.LoginRequest) (string, error) {

	user, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}
		return "", err
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	
	claims := jwt.MapClaims{
		"sub": user.Username, 
		"exp": time.Now().Add(u.tokenTTL).Unix(), 
		"iat": time.Now().Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	
	tokenString, err := token.SignedString(u.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}