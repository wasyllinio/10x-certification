-- Migration: Create Locations Table
-- Created: 2025-10-20 23:03:03
-- Description: Create locations table for EV charging station locations

CREATE TABLE locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    country_code CHAR(3) NOT NULL CHECK (country_code ~ '^[A-Z]{3}$'),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- Add comments for documentation
COMMENT ON TABLE locations IS 'Physical locations where EV charging stations can be installed';
COMMENT ON COLUMN locations.country_code IS 'ISO 3166-1 alpha-3 country code (3 uppercase letters)';
COMMENT ON COLUMN locations.version IS 'Version for optimistic locking';
COMMENT ON COLUMN locations.deleted_at IS 'Soft delete timestamp - NULL means active location';
