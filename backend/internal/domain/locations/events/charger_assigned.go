package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ChargerAssigned represents the event when a charger is assigned to a location
type ChargerAssigned struct {
	events.BaseEvent
	ChargerID  uuid.UUID
	LocationID uuid.UUID
	OwnerID    uuid.UUID
	Timestamp  time.Time
}

// NewChargerAssigned creates a new ChargerAssigned event
func NewChargerAssigned(chargerID, locationID, ownerID uuid.UUID) *ChargerAssigned {
	return &ChargerAssigned{
		BaseEvent:  events.NewBaseEvent("ChargerAssigned"),
		ChargerID:  chargerID,
		LocationID: locationID,
		OwnerID:    ownerID,
		Timestamp:  time.Now(),
	}
}
