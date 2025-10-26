package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LocationDB represents the locations table in the database
type LocationDB struct {
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamptz;not null;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index"`
	Name        string         `gorm:"column:name;type:varchar(255);not null"`
	Address     string         `gorm:"column:address;type:text;not null"`
	CountryCode string         `gorm:"column:country_code;type:char(3);not null;check:country_code ~ '^[A-Z]{3}$'"`
	Chargers    []ChargerDB    `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:SET NULL"`
	EVSEs       []EVSEDB       `gorm:"foreignKey:LocationID;references:ID;constraint:OnDelete:CASCADE"`
	Owner       UserDB         `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT"`
	Version     int            `gorm:"column:version;type:integer;not null;default:1"`
	ID          uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	OwnerID     uuid.UUID      `gorm:"column:owner_id;type:uuid;not null;index:idx_locations_owner,where:deleted_at IS NULL"`
}

// TableName returns the table name for LocationDB
func (LocationDB) TableName() string {
	return "locations"
}

// NewLocationDB creates a new LocationDB with generated ID
func NewLocationDB() *LocationDB {
	return &LocationDB{
		ID: uuid.New(),
	}
}

// BeforeCreate sets up indexes for pagination and search
func (l *LocationDB) BeforeCreate(tx *gorm.DB) error {
	// Create pagination index
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_locations_pagination 
		ON locations(created_at DESC, id) 
		WHERE deleted_at IS NULL
	`)

	// Create search index using GIN for full-text search
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_locations_search 
		ON locations USING GIN (name gin_trgm_ops, address gin_trgm_ops)
	`)

	return nil
}
