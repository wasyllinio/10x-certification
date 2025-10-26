package repository

import (
	"10x-certification/internal/infrastructure/persistence/models"
	"context"

	"github.com/google/uuid"
)

// UserRepository defines the interface for user repository
type UserRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*models.UserDB, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.UserDB, error)
	Save(ctx context.Context, user *models.UserDB) error
}
