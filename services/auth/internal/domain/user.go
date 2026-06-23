package domain

import "time"

type User struct {
	ID           int64
	Username     string
	PasswordHash string
	Name         string
	Surname      string
	Email        string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Surname  string `json:"surname" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=72"`
}
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
}
