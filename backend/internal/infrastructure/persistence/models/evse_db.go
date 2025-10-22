package models

import (
	"time"

	"gorm.io/gorm"
	"github.com/google/uuid"
)

// EVSEDB represents the evse table in the database
type EVSEDB struct {
	ID          uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	EvseID      string         `gorm:"column:evse_id;type:varchar(50);not null;uniqueIndex:idx_evse_id_unique,where:deleted_at IS NULL;check:evse_id ~ '^[A-Z]{2}\\*[A-Z0-9]{3}\\*E[A-Z0-9\\*]+$'"`
	ConnectorID uuid.UUID      `gorm:"column:connector_id;type:uuid;not null;unique;index:idx_evse_connector,where:deleted_at IS NULL"`
	LocationID  uuid.UUID      `gorm:"column:location_id;type:uuid;not null;index:idx_evse_location,where:deleted_at IS NULL"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index"`

	// Relationships
	Connector ConnectorDB `gorm:"foreignKey:ConnectorID;references:ID;constraint:OnDelete:CASCADE"`
	Location  LocationDB  `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:CASCADE"`
}

// TableName returns the table name for EVSEDB
func (EVSEDB) TableName() string {
	return "evse"
}

// BeforeCreate sets up pagination index
func (e *EVSEDB) BeforeCreate(tx *gorm.DB) error {
	// Create pagination index
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_evse_pagination 
		ON evse(created_at DESC, id) 
		WHERE deleted_at IS NULL
	`)

	return nil
}
