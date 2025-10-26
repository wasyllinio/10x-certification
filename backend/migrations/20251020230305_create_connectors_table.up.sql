-- Migration: Create Connectors Table
-- Created: 2025-10-20 23:03:05
-- Description: Create connectors table for EV charger ports

CREATE TABLE connectors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    charger_id UUID NOT NULL REFERENCES chargers(id) ON DELETE CASCADE,
    connector_id INTEGER NOT NULL CHECK (connector_id > 0),
    power DECIMAL(9,1) NOT NULL CHECK (power > 0),
    voltage INTEGER NOT NULL CHECK (voltage > 0),
    amperage INTEGER NOT NULL CHECK (amperage > 0),
    connector_type connector_type NOT NULL,
    connector_standard connector_standard NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT uq_charger_connector UNIQUE (charger_id, connector_id)
);

-- Add comments for documentation
COMMENT ON TABLE connectors IS 'Physical connector ports on EV charging stations';
COMMENT ON COLUMN connectors.charger_id IS 'Reference to the charging station';
COMMENT ON COLUMN connectors.connector_id IS 'Connector ID within the charging station (must be unique per charger)';
COMMENT ON COLUMN connectors.power IS 'Power output in kW (e.g., 22.0, 50.0)';
COMMENT ON COLUMN connectors.voltage IS 'Voltage in volts (e.g., 230, 400)';
COMMENT ON COLUMN connectors.amperage IS 'Amperage in amperes (e.g., 16, 32)';
COMMENT ON COLUMN connectors.connector_type IS 'Type of connector (CCS, Type2, Chademo)';
COMMENT ON COLUMN connectors.connector_standard IS 'Connector standard (AC_1P, AC_3P, DC)';
COMMENT ON COLUMN connectors.deleted_at IS 'Soft delete timestamp - NULL means active connector';
