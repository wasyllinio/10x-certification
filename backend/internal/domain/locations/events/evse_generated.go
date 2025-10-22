package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// EVSEGenerated represents the event when EVSE points are generated
type EVSEGenerated struct {
	events.BaseEvent
	EVSEID      uuid.UUID
	ConnectorID uuid.UUID
	LocationID  uuid.UUID
	EvseID      string
	Timestamp   time.Time
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
