package repository

import (
	"10x-certification/internal/infrastructure/persistence/models"
	"context"
)

// ChargerRepository defines the interface for charger repository
type ChargerRepository interface {
	Create(ctx context.Context, chargerDB *models.ChargerDB) error
	ExistsByVendorAndSerial(ctx context.Context, vendor, serialNumber string) (bool, error)
}
