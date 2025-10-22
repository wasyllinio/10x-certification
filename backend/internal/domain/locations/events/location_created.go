package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// LocationCreated represents the event when a location is created
type LocationCreated struct {
	events.BaseEvent
	LocationID  uuid.UUID
	Name        string
	Address     string
	CountryCode string
	OwnerID     uuid.UUID
	Timestamp   time.Time
}

// NewLocationCreated creates a new LocationCreated event
func NewLocationCreated(locationID, ownerID uuid.UUID, name, address, countryCode string) *LocationCreated {
	return &LocationCreated{
		BaseEvent:   events.NewBaseEvent("LocationCreated"),
		LocationID:  locationID,
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		OwnerID:     ownerID,
		Timestamp:   time.Now(),
	}
}
