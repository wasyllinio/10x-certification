package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ChargerCreated represents the event when a charger is created
type ChargerCreated struct {
	Timestamp time.Time
	events.BaseEvent
	Vendor       string
	Model        string
	SerialNumber string
	ChargerID    uuid.UUID
	OwnerID      uuid.UUID
}

// NewChargerCreated creates a new ChargerCreated event
func NewChargerCreated(chargerID, ownerID uuid.UUID, vendor, model, serialNumber string) *ChargerCreated {
	return &ChargerCreated{
		BaseEvent:    events.NewBaseEvent("ChargerCreated"),
		ChargerID:    chargerID,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		Timestamp:    time.Now(),
	}
}
