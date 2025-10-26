package response

import (
	"time"

	"github.com/google/uuid"
)

// ChargerListItemResponse represents a charger item in a list response
type ChargerListItemResponse struct {
	CreatedAt    time.Time                           `json:"created_at"`
	LocationID   *uuid.UUID                          `json:"location_id"`
	Vendor       string                              `json:"vendor"`
	Model        string                              `json:"model"`
	SerialNumber string                              `json:"serial_number"`
	Status       string                              `json:"status"`
	Connectors   []*ChargerListItemConnectorResponse `json:"connectors"`
	ID           uuid.UUID                           `json:"id"`
	OwnerID      uuid.UUID                           `json:"owner_id"`
}

// NewChargerListItemResponse creates a new ChargerListItemResponse from domain Charger
func NewChargerListItemResponse(
	id uuid.UUID,
	vendor, model, serialNumber, status string,
	ownerID uuid.UUID,
	locationID *uuid.UUID,
	connectors []*ChargerListItemConnectorResponse,
	createdAt time.Time,
) *ChargerListItemResponse {
	return &ChargerListItemResponse{
		ID:           id,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		Status:       status,
		LocationID:   locationID,
		Connectors:   connectors,
		CreatedAt:    createdAt,
	}
}
