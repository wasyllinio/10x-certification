-- Migration: Create EVSE Table
-- Created: 2025-10-20 23:03:06
-- Description: Create EVSE table for EV charging points

CREATE TABLE evse (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    evse_id VARCHAR(50) NOT NULL CHECK (evse_id ~ '^[A-Z]{2}\*[A-Z0-9]{3}\*E[A-Z0-9\*]+$'),
    connector_id UUID UNIQUE NOT NULL REFERENCES connectors(id) ON DELETE CASCADE,
    location_id UUID NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);

-- Add comments for documentation
COMMENT ON TABLE evse IS 'EV charging points generated from connectors at locations';
COMMENT ON COLUMN evse.evse_id IS 'EVSE ID in Emi3spec format (e.g., PL*123*E123*1)';
COMMENT ON COLUMN evse.connector_id IS 'Unique reference to connector (1:1 relationship)';
COMMENT ON COLUMN evse.location_id IS 'Location where this EVSE is available';
COMMENT ON COLUMN evse.deleted_at IS 'Soft delete timestamp - NULL means active EVSE';
