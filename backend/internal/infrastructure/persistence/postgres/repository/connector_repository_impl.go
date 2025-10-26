package repository

import (
	"10x-certification/internal/infrastructure/persistence/models"
	"context"

	"gorm.io/gorm"
)

// ConnectorRepositoryImpl implements ConnectorRepository interface using PostgreSQL
type ConnectorRepositoryImpl struct {
	db *gorm.DB
}

// NewConnectorRepository creates a new ConnectorRepository implementation
func NewConnectorRepository(db interface{}) *ConnectorRepositoryImpl {
	return &ConnectorRepositoryImpl{
		db: db.(*gorm.DB),
	}
}

// Create creates a connector
func (r *ConnectorRepositoryImpl) Create(ctx context.Context, connectorDB *models.ConnectorDB) error {
	return r.db.WithContext(ctx).Create(connectorDB).Error
}
