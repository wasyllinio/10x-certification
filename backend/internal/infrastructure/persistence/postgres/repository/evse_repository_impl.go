package repository

import (
	"gorm.io/gorm"
)

// EVSERepositoryImpl implements EVSERepository interface using PostgreSQL
type EVSERepositoryImpl struct {
	db *gorm.DB
}

// NewLocationRepository creates a new LocationRepository implementation
func NewEVSERepository(db interface{}) *EVSERepositoryImpl {
	return &EVSERepositoryImpl{
		db: db.(*gorm.DB),
	}
}
