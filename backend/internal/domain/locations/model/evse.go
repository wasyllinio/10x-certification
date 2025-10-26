package model

import (
	"time"

	"github.com/google/uuid"
)

// EVSE represents an Electric Vehicle Supply Equipment entity
type EVSE struct {
	CreatedAt   time.Time
	EvseID      string
	ID          uuid.UUID
	ConnectorID uuid.UUID
	LocationID  uuid.UUID
}

// NewEVSE creates a new EVSE entity
func NewEVSE(evseID string, connectorID, locationID uuid.UUID) *EVSE {
	return &EVSE{
		ID:          uuid.New(),
		EvseID:      evseID,
		ConnectorID: connectorID,
		LocationID:  locationID,
		CreatedAt:   time.Now(),
	}
}

// ValidateEvseID validates EVSE ID format according to Emi3spec
func ValidateEvseID(evseID string) error {
	// TODO: Implement EVSE ID validation according to Emi3spec format
	// Format: ^[A-Z]{2}\*[A-Z0-9]{3}\*E[A-Z0-9\*]+$
	panic("not implemented")
}
