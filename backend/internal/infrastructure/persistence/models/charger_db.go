package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ChargerDB represents the chargers table in the database
type ChargerDB struct {
	CreatedAt            time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	UpdatedAt            time.Time      `gorm:"column:updated_at;type:timestamptz;not null;default:now()"`
	AssignedToLocationAt *time.Time     `gorm:"column:assigned_to_location_at;type:timestamptz"`
	LocationID           *uuid.UUID     `gorm:"column:location_id;type:uuid;index:idx_chargers_location,where:deleted_at IS NULL"`
	LastStatusChangeAt   *time.Time     `gorm:"column:last_status_change_at;type:timestamptz"`
	Location             *LocationDB    `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:SET NULL"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index"`
	SerialNumber         string         `gorm:"column:serial_number;type:varchar(255);not null"`
	Model                string         `gorm:"column:model;type:varchar(255);not null"`
	Vendor               string         `gorm:"column:vendor;type:varchar(255);not null"`
	Connectors           []ConnectorDB  `gorm:"foreignKey:ChargerID;references:ID;constraint:OnDelete:CASCADE"`
	Owner                UserDB         `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT"`
	Version              int            `gorm:"column:version;type:integer;not null;default:1"`
	OwnerID              uuid.UUID      `gorm:"column:owner_id;type:uuid;not null;index:idx_chargers_owner,where:deleted_at IS NULL"`
	ID                   uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
}

// TableName returns the table name for ChargerDB
func (ChargerDB) TableName() string {
	return "chargers"
}

// NewChargerDB creates a new ChargerDB with generated ID
func NewChargerDB() *ChargerDB {
	return &ChargerDB{
		ID: uuid.New(),
	}
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
