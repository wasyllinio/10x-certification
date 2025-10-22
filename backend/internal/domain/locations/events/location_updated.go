package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// LocationUpdated represents the event when a location is updated
type LocationUpdated struct {
	events.BaseEvent
	LocationID  uuid.UUID
	Name        string
	Address     string
	CountryCode string
	OwnerID     uuid.UUID
	Timestamp   time.Time
}

// NewLocationUpdated creates a new LocationUpdated event
func NewLocationUpdated(locationID, ownerID uuid.UUID, name, address, countryCode string) *LocationUpdated {
	return &LocationUpdated{
		BaseEvent:   events.NewBaseEvent("LocationUpdated"),
		LocationID:  locationID,
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		OwnerID:     ownerID,
		Timestamp:   time.Now(),
	}
}
