package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ConnectorType represents connector type enum for database
type ConnectorType string

const (
	ConnectorTypeCCS     ConnectorType = "CCS"
	ConnectorTypeType2   ConnectorType = "Type2"
	ConnectorTypeChademo ConnectorType = "Chademo"
)

// ConnectorStandard represents connector standard enum for database
type ConnectorStandard string

const (
	ConnectorStandardAC1P ConnectorStandard = "AC_1P"
	ConnectorStandardAC3P ConnectorStandard = "AC_3P"
	ConnectorStandardDC   ConnectorStandard = "DC"
)

// ConnectorDB represents the connectors table in the database
type ConnectorDB struct {
	CreatedAt         time.Time         `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	UpdatedAt         time.Time         `gorm:"column:updated_at;type:timestamptz;not null;default:now()"`
	EVSE              *EVSEDB           `gorm:"foreignKey:ConnectorID;references:ID;constraint:OnDelete:CASCADE"`
	DeletedAt         gorm.DeletedAt    `gorm:"column:deleted_at;type:timestamptz;index"`
	ConnectorType     ConnectorType     `gorm:"column:connector_type;type:connector_type;not null"`
	ConnectorStandard ConnectorStandard `gorm:"column:connector_standard;type:connector_standard;not null"`
	Charger           ChargerDB         `gorm:"foreignKey:ChargerID;references:ID;constraint:OnDelete:CASCADE"`
	Amperage          int               `gorm:"column:amperage;type:integer;not null;check:amperage > 0"`
	Voltage           int               `gorm:"column:voltage;type:integer;not null;check:voltage > 0"`
	Power             float64           `gorm:"column:power;type:decimal(9,1);not null;check:power > 0"`
	ConnectorID       int               `gorm:"column:connector_id;type:integer;not null;check:connector_id > 0"`
	ID                uuid.UUID         `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	ChargerID         uuid.UUID         `gorm:"column:charger_id;type:uuid;not null;index:idx_connectors_charger,where:deleted_at IS NULL"`
}

// TableName returns the table name for ConnectorDB
func (ConnectorDB) TableName() string {
	return "connectors"
}

// NewConnectorDB creates a new ConnectorDB with generated ID
func NewConnectorDB() *ConnectorDB {
	return &ConnectorDB{
		ID: uuid.New(),
	}
}

// BeforeCreate sets up unique constraint and indexes
func (c *ConnectorDB) BeforeCreate(tx *gorm.DB) error {
	// Create unique constraint for charger_id + connector_id with soft delete
	tx.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_connectors_unique 
		ON connectors(charger_id, connector_id) 
		WHERE deleted_at IS NULL
	`)

	// Create pagination index
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_connectors_pagination 
		ON connectors(created_at DESC, id) 
		WHERE deleted_at IS NULL
	`)

	return nil
}
