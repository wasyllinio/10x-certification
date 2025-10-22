package response

import (
	"time"

	"github.com/google/uuid"
)

// ChargerDetailResponse represents a detailed charger response DTO
type ChargerDetailResponse struct {
	ID           uuid.UUID           `json:"id"`
	Vendor       string              `json:"vendor"`
	Model        string              `json:"model"`
	SerialNumber string              `json:"serial_number"`
	OwnerID      uuid.UUID           `json:"owner_id"`
	LocationID   *uuid.UUID          `json:"location_id,omitempty"`
	Status       string              `json:"status"`
	Connectors   []ConnectorResponse `json:"connectors"`
	CreatedAt    time.Time           `json:"created_at"`
}

// NewChargerDetailResponse creates a new ChargerDetailResponse from domain Charger
func NewChargerDetailResponse(
	id, ownerID uuid.UUID,
	vendor, model, serialNumber, status string,
	locationID *uuid.UUID,
	connectors []ConnectorResponse,
	createdAt time.Time,
) *ChargerDetailResponse {
	return &ChargerDetailResponse{
		ID:           id,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		LocationID:   locationID,
		Status:       status,
		Connectors:   connectors,
		CreatedAt:    createdAt,
	}
}
