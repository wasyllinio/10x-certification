package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ChargerDetached represents the event when a charger is detached from a location
type ChargerDetached struct {
	Timestamp time.Time
	events.BaseEvent
	ChargerID  uuid.UUID
	LocationID uuid.UUID
	OwnerID    uuid.UUID
}

// NewChargerDetached creates a new ChargerDetached event
func NewChargerDetached(chargerID, locationID, ownerID uuid.UUID) *ChargerDetached {
	return &ChargerDetached{
		BaseEvent:  events.NewBaseEvent("ChargerDetached"),
		ChargerID:  chargerID,
		LocationID: locationID,
		OwnerID:    ownerID,
		Timestamp:  time.Now(),
	}
}
