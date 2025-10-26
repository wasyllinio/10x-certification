-- Migration: Drop Additional Composite Indexes
-- Created: 2025-10-20 23:03:09
-- Description: Drop additional composite indexes (reverse of 20251020230309_create_additional_indexes_up.sql)

-- Drop additional search optimization indexes
DROP INDEX IF EXISTS idx_chargers_model_lower;
DROP INDEX IF EXISTS idx_chargers_vendor_lower;
DROP INDEX IF EXISTS idx_locations_name_lower;
DROP INDEX IF EXISTS idx_users_email_lower;

-- Drop composite indexes for common queries
DROP INDEX IF EXISTS idx_chargers_owner_location;
