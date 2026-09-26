package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/bekontaii/Booking_System_CyberSlotGin/services/internal/domain"
)

type MemoryClubRepository struct {
	clubs  map[int64]*domain.Club
	nextID int64
	mu     sync.RWMutex
}

func NewMemoryClubRepository() *MemoryClubRepository {
	return &MemoryClubRepository{
		clubs:  make(map[int64]*domain.Club),
		nextID: 1,
	}
}
func (r *MemoryClubRepository) CreateClub(ctx context.Context, club *domain.Club) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	club.ID = r.nextID
	now := time.Now()
	club.CreatedAt = now
	club.UpdatedAt = now
	r.clubs[club.ID] = club
	r.nextID++
	return nil
}
func (r *MemoryClubRepository) GetClubByID(ctx context.Context, clubID int64) (*domain.Club,error){
	r.mu.RLock()
	defer r.mu.RUnlock()
	club, ok := r.clubs[clubID]
	if !ok {
		return nil, errors.New("club not found")
	}
	return club,nil
}

func (r *MemoryClubRepository) GetAllClubs(ctx context.Context) ([]domain.Club,error){
	r.mu.RLock()
	defer r.mu.RUnlock()
	clubs :=make([]domain.Club,0,len(r.clubs))
	for _,club := range r.clubs{
		clubs= append(clubs, *club)
	}
	return clubs,nil
}
func (r *MemoryClubRepository) UpdateClub(ctx context.Context,club *domain.Club) error{
	r.mu.Lock()
	defer r.mu.Unlock()
	existingClub,ok := r.clubs[club.ID]
		if !ok {
			return errors.New("club not found")
		}
		createdAt := existingClub.CreatedAt
		*existingClub = *club
		existingClub.CreatedAt = createdAt
		existingClub.UpdatedAt = time.Now()
		return nil
}
func (r *MemoryClubRepository) DeleteClub(ctx context.Context,clubID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	foundClub,ok := r.clubs[clubID]
	if !ok{
		return errors.New("club not found")
	}
	delete(r.clubs,foundClub.ID)
	return nil

}