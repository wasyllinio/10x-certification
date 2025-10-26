package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ConnectorAdded represents the event when a connector is added to a charger
type ConnectorAdded struct {
	Timestamp time.Time
	events.BaseEvent
	ConnectorType     string
	ConnectorStandard string
	ConnectorIDInt    int
	Voltage           int
	Amperage          int
	Power             float32
	ConnectorID       uuid.UUID
	ChargerID         uuid.UUID
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
