package repository

import (
	"10x-certification/internal/infrastructure/persistence/models"
	"context"
)

// ConnectorRepository defines the interface for connector repository
type ConnectorRepository interface {
	Create(ctx context.Context, connectorDB *models.ConnectorDB) error
}
