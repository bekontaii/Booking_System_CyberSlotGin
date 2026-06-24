package repository

import (
	"context"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"sync"
)

type MemoryUserRepository struct {
	users  map[int64]*domain.User
	mu     sync.RWMutex
	nextID int64
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int64]*domain.User),
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

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}
func (r *MemoryUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}
func (r *MemoryUserRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	value, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return value, nil
}
