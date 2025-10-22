package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ConnectorAdded represents the event when a connector is added to a charger
type ConnectorAdded struct {
	events.BaseEvent
	ConnectorID       uuid.UUID
	ChargerID         uuid.UUID
	ConnectorIDInt    int
	Power             float32
	Voltage           int
	Amperage          int
	ConnectorType     string
	ConnectorStandard string
	Timestamp         time.Time
}

// NewConnectorAdded creates a new ConnectorAdded event
func NewConnectorAdded(
	connectorID, chargerID uuid.UUID,
	connectorIDInt int,
	power float32,
	voltage, amperage int,
	connectorType, connectorStandard string,
) *ConnectorAdded {
	return &ConnectorAdded{
		BaseEvent:         events.NewBaseEvent("ConnectorAdded"),
		ConnectorID:       connectorID,
		ChargerID:         chargerID,
		ConnectorIDInt:    connectorIDInt,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
		Timestamp:         time.Now(),
	}
}
