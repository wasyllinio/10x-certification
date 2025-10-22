-- Migration: Create Additional Composite Indexes
-- Created: 2025-10-20 23:03:09
-- Description: Create composite indexes for owner+location queries and search optimization

-- Composite indexes for common queries
CREATE INDEX idx_chargers_owner_location ON chargers(owner_id, location_id) WHERE deleted_at IS NULL;

-- Additional search optimization indexes
CREATE INDEX idx_users_email_lower ON users(LOWER(email)) WHERE deleted_at IS NULL;
CREATE INDEX idx_locations_name_lower ON locations(LOWER(name)) WHERE deleted_at IS NULL;
CREATE INDEX idx_chargers_vendor_lower ON chargers(LOWER(vendor)) WHERE deleted_at IS NULL;
CREATE INDEX idx_chargers_model_lower ON chargers(LOWER(model)) WHERE deleted_at IS NULL;
