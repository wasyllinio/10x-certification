-- Migration: Drop Custom ENUM Types
-- Created: 2025-10-20 23:03:01
-- Description: Drop custom ENUM types (reverse of 20251020230301_create_custom_types_up.sql)

-- Drop audit operation enum
DROP TYPE IF EXISTS audit_operation;

-- Drop connector standard enum
DROP TYPE IF EXISTS connector_standard;

-- Drop connector type enum
DROP TYPE IF EXISTS connector_type;
