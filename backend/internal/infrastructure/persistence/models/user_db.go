package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole represents user role enum for database
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleOwner UserRole = "owner"
)

// UserDB represents the users table in the database
type UserDB struct {
	CreatedAt       time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now()"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;type:timestamptz;not null;default:now()"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at;type:timestamptz;index"`
	Email           string         `gorm:"column:email;type:varchar(255);not null;uniqueIndex:idx_users_email,where:deleted_at IS NULL"`
	PasswordHash    string         `gorm:"column:password_hash;type:varchar(255);not null"`
	PasswordSalt    string         `gorm:"column:password_salt;type:varchar(255);not null"`
	Role            UserRole       `gorm:"column:role;type:varchar(20);not null;check:role IN ('admin', 'owner')"`
	AuthorizationID uuid.UUID         `gorm:"column:authorization_id;type:uuid;not null"`
	Chargers        []ChargerDB    `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT"`
	Locations       []LocationDB   `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:RESTRICT"`
	AuditLogs       []AuditLogDB   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:RESTRICT"`
	ID              uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
}

// TableName returns the table name for UserDB
func (UserDB) TableName() string {
	return "users"
}

// NewUserDB creates a new UserDB with generated ID
func NewUserDB() *UserDB {
	return &UserDB{
		ID: uuid.New(),
		AuthorizationID: uuid.New(),
	}
}
