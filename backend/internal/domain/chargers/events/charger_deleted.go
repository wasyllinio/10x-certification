package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// ChargerDeleted represents the event when a charger is deleted
type ChargerDeleted struct {
	Timestamp time.Time
	events.BaseEvent
	Vendor       string
	Model        string
	SerialNumber string
	ChargerID    uuid.UUID
	OwnerID      uuid.UUID
}

// NewChargerDeleted creates a new ChargerDeleted event
func NewChargerDeleted(chargerID, ownerID uuid.UUID, vendor, model, serialNumber string) *ChargerDeleted {
	return &ChargerDeleted{
		BaseEvent:    events.NewBaseEvent("ChargerDeleted"),
		ChargerID:    chargerID,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		Timestamp:    time.Now(),
	}
}
