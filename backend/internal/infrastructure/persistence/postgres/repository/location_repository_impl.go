package repository

import (
	"gorm.io/gorm"
)

// LocationRepositoryImpl implements LocationRepository interface using PostgreSQL
type LocationRepositoryImpl struct {
	db *gorm.DB
}

// NewLocationRepository creates a new LocationRepository implementation
func NewLocationRepository(db interface{}) *LocationRepositoryImpl {
	return &LocationRepositoryImpl{
		db: db.(*gorm.DB),
	}
}
