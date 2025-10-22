-- Migration: Drop Performance Indexes
-- Created: 2025-10-20 23:03:08
-- Description: Drop all performance indexes (reverse of 20251020230308_create_performance_indexes_up.sql)

-- Drop audit logs indexes
DROP INDEX IF EXISTS idx_audit_jsonb;
DROP INDEX IF EXISTS idx_audit_operation_time;
DROP INDEX IF EXISTS idx_audit_table_record;
DROP INDEX IF EXISTS idx_audit_user_time;

-- Drop GIN indexes for full-text search
DROP INDEX IF EXISTS idx_locations_search;
DROP INDEX IF EXISTS idx_chargers_search;

-- Drop pagination indexes
DROP INDEX IF EXISTS idx_evse_pagination;
DROP INDEX IF EXISTS idx_connectors_pagination;
DROP INDEX IF EXISTS idx_locations_pagination;
DROP INDEX IF EXISTS idx_chargers_pagination;

-- Drop foreign key indexes
DROP INDEX IF EXISTS idx_evse_location;
DROP INDEX IF EXISTS idx_evse_connector;
DROP INDEX IF EXISTS idx_connectors_charger;
DROP INDEX IF EXISTS idx_locations_owner;
DROP INDEX IF EXISTS idx_chargers_location;
DROP INDEX IF EXISTS idx_chargers_owner;

-- Drop unique indexes
DROP INDEX IF EXISTS idx_evse_id_unique;
DROP INDEX IF EXISTS idx_connectors_unique;
DROP INDEX IF EXISTS idx_chargers_serial;
DROP INDEX IF EXISTS idx_users_email;
