package response

import (
	"time"

	"github.com/google/uuid"
)

// ChargerListItemResponse represents a charger item in a list response
type ChargerListItemResponse struct {
	ID           uuid.UUID                           `json:"id"`
	Vendor       string                              `json:"vendor"`
	Model        string                              `json:"model"`
	SerialNumber string                              `json:"serial_number"`
	OwnerID      uuid.UUID                           `json:"owner_id"`
	Status       string                              `json:"status"`
	LocationID   *uuid.UUID                          `json:"location_id"`
	Connectors   []*ChargerListItemConnectorResponse `json:"connectors"`
	CreatedAt    time.Time                           `json:"created_at"`
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
