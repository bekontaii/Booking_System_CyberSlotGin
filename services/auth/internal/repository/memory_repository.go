package repository

import (
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"sync"
	"context"
)

type MemoryUserRepository struct {
	users map[int64]*domain.User
	mu   sync.RWMutex
	nextID int64
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[int64]*domain.User),
		nextID: 1,
	}
}
func (r *MemoryUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return nil
}
func (r *MemoryUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()