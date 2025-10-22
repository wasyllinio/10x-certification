package response

import (
	"time"

	"github.com/google/uuid"
)

// ChargerListItemResponse represents a charger item in a list response
type ChargerListItemResponse struct {
	ID           uuid.UUID  `json:"id"`
	Vendor       string     `json:"vendor"`
	Model        string     `json:"model"`
	SerialNumber string     `json:"serial_number"`
	Status       string     `json:"status"`
	LocationID   *uuid.UUID `json:"location_id,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

// NewChargerListItemResponse creates a new ChargerListItemResponse from domain Charger
func NewChargerListItemResponse(
	id uuid.UUID,
	vendor, model, serialNumber, status string,
	locationID *uuid.UUID,
	createdAt time.Time,
) *ChargerListItemResponse {
	return &ChargerListItemResponse{
		ID:           id,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		Status:       status,
		LocationID:   locationID,
		CreatedAt:    createdAt,
	}
}
