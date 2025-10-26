package repository

import (
	"10x-certification/internal/infrastructure/persistence/models"
	"context"

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

// Create creates a charger with connectors in a transaction
func (r *ChargerRepositoryImpl) Create(ctx context.Context, chargerDB *models.ChargerDB, connectorsDB []models.ConnectorDB) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create charger
		if err := tx.Create(chargerDB).Error; err != nil {
			return err
		}

		// Create connectors if any
		if len(connectorsDB) > 0 {
			if err := tx.Create(&connectorsDB).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// ExistsByVendorAndSerial checks if a charger with given vendor and serial number exists
func (r *ChargerRepositoryImpl) ExistsByVendorAndSerial(ctx context.Context, vendor, serialNumber string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.ChargerDB{}).
		Where("vendor = ? AND serial_number = ? AND deleted_at IS NULL", vendor, serialNumber).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
