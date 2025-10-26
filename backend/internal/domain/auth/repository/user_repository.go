package repository

import (
	"10x-certification/internal/infrastructure/persistence/models"
	"context"
)

// UserRepository defines the interface for user repository
type UserRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Save(ctx context.Context, user *models.UserDB) error
}
