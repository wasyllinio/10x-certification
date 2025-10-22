package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

// ChargerDB represents the chargers table in the database
type ChargerDB struct {
	ID                   string         `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	Vendor               string         `gorm:"column:vendor;type:varchar(255);not null"`
	Model                string         `gorm:"column:model;type:varchar(255);not null"`
	SerialNumber         string         `gorm:"column:serial_number;type:varchar(255);not null"`
	OwnerID              uuid.UUID      `gorm:"column:owner_id;type:uuid;not null;index:idx_chargers_owner,where:deleted_at IS NULL"`
	LocationID           *uuid.UUID     `gorm:"column:location_id;type:uuid;index:idx_chargers_location,where:deleted_at IS NULL"`
	AssignedToLocationAt *time.Time     `gorm:"column:assigned_to_location_at;type:timestamptz"`
	LastStatusChangeAt   *time.Time     `gorm:"column:last_status_change_at;type:timestamptz"`
	Version              int            `gorm:"column:version;type:integer;not null;default:1"`
	CreatedAt            time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	UpdatedAt            time.Time      `gorm:"column:updated_at;type:timestamptz;not null;default:now()"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index"`

	// Relationships
	Owner      UserDB        `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT"`
	Location   *LocationDB   `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:SET NULL"`
	Connectors []ConnectorDB `gorm:"foreignKey:ChargerID;references:ID;constraint:OnDelete:CASCADE"`
}

// TableName returns the table name for ChargerDB
func (ChargerDB) TableName() string {
	return "chargers"
}

// BeforeCreate sets up unique index constraint for vendor + serial_number
func (c *ChargerDB) BeforeCreate(tx *gorm.DB) error {
	// Create unique index for vendor + serial_number with soft delete condition
	tx.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_chargers_serial 
		ON chargers(vendor, serial_number) 
		WHERE deleted_at IS NULL
	`)
	return nil
}
