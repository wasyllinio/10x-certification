package repository

import (
	"gorm.io/gorm"
)

// ConnectorRepositoryImpl implements ConnectorRepository interface using PostgreSQL
type ConnectorRepositoryImpl struct {
	db *gorm.DB
}

// NewChargerRepository creates a new ChargerRepository implementation
func NewConnectorRepository(db interface{}) *ConnectorRepositoryImpl {
	return &ConnectorRepositoryImpl{
		db: db.(*gorm.DB),
	}
}
