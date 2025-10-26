package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ConnectorRemoved represents the event when a connector is removed from a charger
type ConnectorRemoved struct {
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

// NewConnectorRemoved creates a new ConnectorRemoved event
func NewConnectorRemoved(
	connectorID, chargerID uuid.UUID,
	connectorIDInt int,
	power float32,
	voltage, amperage int,
	connectorType, connectorStandard string,
) *ConnectorRemoved {
	return &ConnectorRemoved{
		BaseEvent:         events.NewBaseEvent("ConnectorRemoved"),
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
