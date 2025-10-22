-- Migration: Create Audit Logs Table
-- Created: 2025-10-20 23:03:07
-- Description: Create audit_logs table for tracking CUD operations (non-partitioned for MVP)

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    operation audit_operation NOT NULL,
    table_name VARCHAR(50) NOT NULL,
    record_id UUID NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    old_values JSONB NULL,
    new_values JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Add comments for documentation
COMMENT ON TABLE audit_logs IS 'Audit trail for all CUD operations (Create, Update, Delete)';
COMMENT ON COLUMN audit_logs.operation IS 'Type of operation performed (INSERT, UPDATE, DELETE)';
COMMENT ON COLUMN audit_logs.table_name IS 'Name of the table that was modified';
COMMENT ON COLUMN audit_logs.record_id IS 'ID of the record that was modified';
COMMENT ON COLUMN audit_logs.user_id IS 'User who performed the operation';
COMMENT ON COLUMN audit_logs.old_values IS 'Previous values (for UPDATE/DELETE operations)';
COMMENT ON COLUMN audit_logs.new_values IS 'New values (for INSERT/UPDATE operations)';
COMMENT ON COLUMN audit_logs.created_at IS 'Timestamp when the operation was performed';
