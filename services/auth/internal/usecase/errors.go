package usecase

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("User already exists")
	ErrEmailAlreadyExists    = errors.New("Email already exists")
	ErrInvalidCredentials    = errors.New("Invalid credentials")
)
