package usecase

import (
	"context"

	"github.com/bekontaii/Booking_System_CyberSlotGin/services/internal/domain"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/internal/repository"
)

type ClubUseCase interface{
	CreateClub(ctx context.Context,club *domain.Club) error
	GetClubByID(ctx context.Context,clubID int64) (*domain.Club,error)
	GetAllClubs(ctx context.Context) ([]domain.Club,error)
	UpdateClub(ctx context.Context,club *domain.Club) error
	DeleteClub(ctx context.Context,clubID int64) error
}
type clubUseCase struct{
	repo repository.ClubRepository
}
func NewClubUseCase (repo repository.ClubRepository) clubUseCase{
	return clubUseCase{
		repo: repo,
	}
}
func (c *clubUseCase) CreateClub(ctx context.Context,club *domain.Club)error {
	err := c.repo.CreateClub(ctx,club)
	if err !=nil {
		return err
	}
	return nil
}
func (c *clubUseCase) GetClubByID(ctx context.Context,clubID int64) (*domain.Club,error){
	club,err:=c.repo.GetClubByID(ctx,clubID)
	if err !=nil{
		return nil,err
	}
	return club,nil
}
func (c *clubUseCase) GetAllClubs(ctx context.Context) ([]domain.Club,error){
	clubs,err:= c.repo.GetAllClubs(ctx)
	if err!=nil{
		return nil,err
	}
	return clubs,nil
}
func (c *clubUseCase) UpdateClub(ctx context.Context,club *domain.Club)error{
	err:=c.repo.UpdateClub(ctx,club)
	if err!=nil{
		return err
	}
	return nil
}
func (c *clubUseCase) DeleteClub(ctx context.Context,clubID int64) error {
	err:=c.repo.DeleteClub(ctx,clubID)
	if err != nil{
		return err
	}
	return nil
}