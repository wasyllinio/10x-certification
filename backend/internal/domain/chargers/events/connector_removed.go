package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ConnectorRemoved represents the event when a connector is removed from a charger
type ConnectorRemoved struct {
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
