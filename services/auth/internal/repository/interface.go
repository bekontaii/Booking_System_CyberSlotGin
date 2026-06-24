package repository
import (
	"github.com/bekontaii/Booking_System_CyberSlotGin/services/auth/internal/domain"
	"context"
)
type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
} 