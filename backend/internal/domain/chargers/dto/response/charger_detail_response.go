package response

import (
	"time"

	"github.com/google/uuid"
)

// ChargerDetailResponse represents a detailed charger response DTO
type ChargerDetailResponse struct {
	CreatedAt    time.Time           `json:"created_at"`
	LocationID   *uuid.UUID          `json:"location_id,omitempty"`
	Vendor       string              `json:"vendor"`
	Model        string              `json:"model"`
	SerialNumber string              `json:"serial_number"`
	Connectors   []ConnectorResponse `json:"connectors"`
}

// NewChargerDetailResponse creates a new ChargerDetailResponse from domain Charger
func NewChargerDetailResponse(
	vendor, model, serialNumber string,
	locationID *uuid.UUID,
	connectors []ConnectorResponse,
	createdAt time.Time,
) *ChargerDetailResponse {
	return &ChargerDetailResponse{
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		LocationID:   locationID,
		Connectors:   connectors,
		CreatedAt:    createdAt,
	}
}
