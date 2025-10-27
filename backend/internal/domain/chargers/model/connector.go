package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// ConnectorType represents connector type enum
type ConnectorType string

const (
	ConnectorTypeCCS     ConnectorType = "CCS"
	ConnectorTypeType2   ConnectorType = "Type2"
	ConnectorTypeChademo ConnectorType = "Chademo"
)

// ConnectorStandard represents connector standard enum
type ConnectorStandard string

const (
	ConnectorStandardAC1P ConnectorStandard = "AC_1P"
	ConnectorStandardAC3P ConnectorStandard = "AC_3P"
	ConnectorStandardDC   ConnectorStandard = "DC"
)

// Connector represents a charging connector entity
type Connector struct {
	CreatedAt         time.Time
	UpdatedAt         time.Time
	ConnectorType     ConnectorType
	ConnectorStandard ConnectorStandard
	ConnectorID       int
	Voltage           int
	Amperage          int
	Power             float32
	ID                uuid.UUID
	ChargerID         uuid.UUID
}

// NewConnector creates a new connector entity
func NewConnector(
	chargerID uuid.UUID,
	connectorID int,
	power float32,
	voltage, amperage int,
	connectorType ConnectorType,
	connectorStandard ConnectorStandard,
) *Connector {
	now := time.Now()
	return &Connector{
		ID:                uuid.New(),
		ChargerID:         chargerID,
		ConnectorID:       connectorID,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
		CreatedAt:         now,
		UpdatedAt:         now,
	}
}

// Validate validates connector data
func (c *Connector) Validate() error {
	if c.ConnectorID <= 0 {
		return errors.New("connector ID must be positive")
	}
	if c.Power <= 0 {
		return errors.New("power must be positive")
	}
	if c.Voltage <= 0 {
		return errors.New("voltage must be positive")
	}
	if c.Amperage <= 0 {
		return errors.New("amperage must be positive")
	}
	return nil
}
