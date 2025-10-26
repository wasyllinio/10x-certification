package repository

import (
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/persistence/models"
	"10x-certification/internal/shared/errors"
	"context"

	"github.com/google/uuid"
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

// FindByEmail finds a user by email
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*models.UserDB, error) {
	var user models.UserDB
	err := r.db.WithContext(ctx).
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// FindByID finds a user by ID
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*models.UserDB, error) {
	var user models.UserDB
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Save saves a user to the database
func (r *UserRepositoryImpl) Save(ctx context.Context, user *models.UserDB) error {
	return r.db.WithContext(ctx).Create(user).Error
}
