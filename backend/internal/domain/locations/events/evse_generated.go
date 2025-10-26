package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// EVSEGenerated represents the event when EVSE points are generated
type EVSEGenerated struct {
	Timestamp time.Time
	events.BaseEvent
	EvseID      string
	EVSEID      uuid.UUID
	ConnectorID uuid.UUID
	LocationID  uuid.UUID
}

// NewEVSEGenerated creates a new EVSEGenerated event
func NewEVSEGenerated(evseID, connectorID, locationID uuid.UUID, evseIDString string) *EVSEGenerated {
	return &EVSEGenerated{
		BaseEvent:   events.NewBaseEvent("EVSEGenerated"),
		EVSEID:      evseID,
		ConnectorID: connectorID,
		LocationID:  locationID,
		EvseID:      evseIDString,
		Timestamp:   time.Now(),
	}
}
