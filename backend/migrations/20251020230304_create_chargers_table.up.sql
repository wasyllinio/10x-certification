-- Migration: Create Chargers Table
-- Created: 2025-10-20 23:03:04
-- Description: Create chargers table for EV charging stations

CREATE TABLE chargers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vendor VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    serial_number VARCHAR(255) NOT NULL,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    location_id UUID NULL REFERENCES locations(id) ON DELETE SET NULL,
    assigned_to_location_at TIMESTAMPTZ NULL,
    last_status_change_at TIMESTAMPTZ NULL,
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- Add comments for documentation
COMMENT ON TABLE chargers IS 'EV charging stations with physical characteristics and ownership';
COMMENT ON COLUMN chargers.location_id IS 'Current location (NULL means in warehouse)';
COMMENT ON COLUMN chargers.assigned_to_location_at IS 'Timestamp when charger was assigned to current location';
COMMENT ON COLUMN chargers.last_status_change_at IS 'Timestamp of last status change';
COMMENT ON COLUMN chargers.version IS 'Version for optimistic locking';
COMMENT ON COLUMN chargers.deleted_at IS 'Soft delete timestamp - NULL means active charger';
