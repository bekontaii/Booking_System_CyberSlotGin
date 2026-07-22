package repository

import (
	"context"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/internal/domain"
)

type ClubRepository interface {
	CreateClub(ctx context.Context, club *domain.Club) error
	GetClub(ctx context.Context, clubID int64) (*domain.Club, error)
	UpdateClub(ctx context.Context, club *domain.Club) error
	DeleteClub(ctx context.Context, clubID string) error
	GetAllClubs() ([]domain.Club, error)
}
