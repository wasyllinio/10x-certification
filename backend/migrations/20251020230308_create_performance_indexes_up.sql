-- Migration: Create Performance Indexes
-- Created: 2025-10-20 23:03:08
-- Description: Create all performance indexes for efficient querying

-- Unique indexes with soft delete conditions
CREATE UNIQUE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_chargers_serial ON chargers(vendor, serial_number) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_connectors_unique ON connectors(charger_id, connector_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_evse_id_unique ON evse(evse_id) WHERE deleted_at IS NULL;

-- Foreign key indexes
CREATE INDEX idx_chargers_owner ON chargers(owner_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_chargers_location ON chargers(location_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_locations_owner ON locations(owner_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_connectors_charger ON connectors(charger_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evse_connector ON evse(connector_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evse_location ON evse(location_id) WHERE deleted_at IS NULL;

-- Pagination indexes (created_at DESC, id)
CREATE INDEX idx_chargers_pagination ON chargers(created_at DESC, id) WHERE deleted_at IS NULL;
CREATE INDEX idx_locations_pagination ON locations(created_at DESC, id) WHERE deleted_at IS NULL;
CREATE INDEX idx_connectors_pagination ON connectors(created_at DESC, id) WHERE deleted_at IS NULL;
CREATE INDEX idx_evse_pagination ON evse(created_at DESC, id) WHERE deleted_at IS NULL;

-- GIN indexes for full-text search
CREATE INDEX idx_chargers_search ON chargers USING GIN (vendor gin_trgm_ops, model gin_trgm_ops, serial_number gin_trgm_ops);
CREATE INDEX idx_locations_search ON locations USING GIN (name gin_trgm_ops, address gin_trgm_ops);

-- Audit logs indexes
CREATE INDEX idx_audit_user_time ON audit_logs(user_id, created_at DESC);
CREATE INDEX idx_audit_table_record ON audit_logs(table_name, record_id, created_at DESC);
CREATE INDEX idx_audit_operation_time ON audit_logs(operation, created_at DESC);
CREATE INDEX idx_audit_jsonb ON audit_logs USING GIN (old_values, new_values);
