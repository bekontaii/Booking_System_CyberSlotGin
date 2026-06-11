package repository
import (
	"context"
	"errors"
	"sync"
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
)
var (
	ErrUserNotFound = errors.New("user not found")
	ErrUsernameExists = errors.New("username already exists")
)
type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}
type InMemoryUserRepository struct {
	mu sync.RWMutex
	users map[string]*domain.User
}
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}
func (r *InMemoryUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.users[user.Username]; exists {
		return ErrUsernameExists
	}
	r.users[user.Username] = user
	return nil
}
func (r *InMemoryUserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, exists := r.users[username]
	if !exists {
		return nil, ErrUserNotFound
	}
	return user, nil
}