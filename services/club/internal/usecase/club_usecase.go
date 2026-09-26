package usecase

import (
	"context"

	"github.com/bekontaii/Booking_System_CyberSlotGin/services/internal/domain"
)

type ClubUseCase interface{
	CreateClub(ctx context.Context,club *domain.Club) error
	GetClubByID(ctx context.Context,clubID int64) (*domain.Club,error)
	GetAllClubs(ctx context.Context) ([]domain.Club,error)
	UpdateClub(ctx context.Context,club *domain.Club) error
	DeleteClub(ctx context.Context,clubID int64) error
}
