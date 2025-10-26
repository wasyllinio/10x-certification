package response

import (
	"github.com/google/uuid"
)

// LocationChargerResponse represents a charger response DTO for locations domain
type LocationChargerResponse struct {
	Vendor       string                      `json:"vendor"`
	Model        string                      `json:"model"`
	SerialNumber string                      `json:"serial_number"`
	Connectors   []LocationConnectorResponse `json:"connectors"`
	ID           uuid.UUID                   `json:"id"`
}

// NewLocationChargerResponse creates a new LocationChargerResponse
func NewLocationChargerResponse(
	id uuid.UUID,
	vendor, model, serialNumber string,
	connectors []LocationConnectorResponse,
) *LocationChargerResponse {
	return &LocationChargerResponse{
		ID:           id,
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		Connectors:   connectors,
	}
}
