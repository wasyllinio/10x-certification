package repository

import (
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/persistence/models"
	"context"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements UserRepository interface using PostgreSQL
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository implementation
func NewUserRepository(db interface{}) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db.(*gorm.DB),
	}
}

// ExistsByEmail checks if a user with given email exists
func (r *UserRepositoryImpl) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.UserDB{}).
		Where("email = ? AND deleted_at IS NULL", email).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Save saves a user to the database
func (r *UserRepositoryImpl) Save(ctx context.Context, user *models.UserDB) error {
	return r.db.WithContext(ctx).Create(user).Error
}
