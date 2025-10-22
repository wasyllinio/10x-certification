package repository

import (
	"gorm.io/gorm"
)

// ChargerRepositoryImpl implements ChargerRepository interface using PostgreSQL
type ChargerRepositoryImpl struct {
	db *gorm.DB
}

// NewChargerRepository creates a new ChargerRepository implementation
func NewChargerRepository(db interface{}) *ChargerRepositoryImpl {
	return &ChargerRepositoryImpl{
		db: db.(*gorm.DB),
	}
}
