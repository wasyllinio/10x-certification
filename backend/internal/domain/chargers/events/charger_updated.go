package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ChargerUpdated represents the event when a charger is updated
type ChargerUpdated struct {
	events.BaseEvent
	ChargerID    uuid.UUID
	Vendor       string
	Model        string
	SerialNumber string
	OwnerID      uuid.UUID
	Timestamp    time.Time
}

// NewChargerUpdated creates a new ChargerUpdated event
func NewChargerUpdated(chargerID, ownerID uuid.UUID, vendor, model, serialNumber string) *ChargerUpdated {
	return &ChargerUpdated{
		BaseEvent:    events.NewBaseEvent("ChargerUpdated"),
		ChargerID:    chargerID,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		Timestamp:    time.Now(),
	}
}
