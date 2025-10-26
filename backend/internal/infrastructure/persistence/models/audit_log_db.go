package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuditOperation represents audit operation enum for database
type AuditOperation string

const (
	AuditOperationInsert AuditOperation = "INSERT"
	AuditOperationUpdate AuditOperation = "UPDATE"
	AuditOperationDelete AuditOperation = "DELETE"
)

// AuditLogDB represents the audit_logs table in the database
// Note: This table uses RANGE partitioning by created_at (monthly) at database level
type AuditLogDB struct {
	CreatedAt      time.Time      `gorm:"column:created_at;type:timestamptz;not null;default:now();index:idx_audit_user_time,priority:2,sort:desc;index:idx_audit_table_record,priority:3,sort:desc;index:idx_audit_operation_time,priority:2,sort:desc"`
	Operation      AuditOperation `gorm:"column:operation;type:audit_operation;not null"`
	TableNameField string         `gorm:"column:table_name;type:varchar(50);not null"`
	OldValues      []byte         `gorm:"column:old_values;type:jsonb"`
	NewValues      []byte         `gorm:"column:new_values;type:jsonb"`
	User           UserDB         `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:RESTRICT"`
	ID             uuid.UUID      `gorm:"column:id;type:uuid;primaryKey;default:gen_random_uuid()"`
	RecordID       uuid.UUID      `gorm:"column:record_id;type:uuid;not null"`
	UserID         uuid.UUID      `gorm:"column:user_id;type:uuid;not null;index:idx_audit_user_time,priority:1"`
}

// TableName returns the table name for AuditLogDB
func (AuditLogDB) TableName() string {
	return "audit_logs"
}

// BeforeCreate sets up additional indexes for efficient querying
func (a *AuditLogDB) BeforeCreate(tx *gorm.DB) error {
	// Create composite index for table_name + record_id + created_at
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_audit_table_record 
		ON audit_logs(table_name, record_id, created_at DESC)
	`)

	// Create composite index for operation + created_at
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_audit_operation_time 
		ON audit_logs(operation, created_at DESC)
	`)

	// Create GIN index for JSONB fields
	tx.Exec(`
		CREATE INDEX IF NOT EXISTS idx_audit_jsonb 
		ON audit_logs USING GIN (old_values, new_values)
	`)

	return nil
}
