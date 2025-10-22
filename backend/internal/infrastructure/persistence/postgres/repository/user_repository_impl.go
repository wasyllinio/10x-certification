package repository

import (
	"10x-certification/internal/domain/auth/repository"

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
