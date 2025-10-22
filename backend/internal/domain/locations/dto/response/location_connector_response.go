package response

import (
	"github.com/google/uuid"
)

// LocationConnectorResponse represents a connector response DTO for locations domain
type LocationConnectorResponse struct {
	ID                uuid.UUID `json:"id"`
	ConnectorID       int       `json:"connector_id"`
	Power             float64   `json:"power"`
	Voltage           int       `json:"voltage"`
	Amperage          int       `json:"amperage"`
	ConnectorType     string    `json:"connector_type"`
	ConnectorStandard string    `json:"connector_standard"`
}

// NewLocationConnectorResponse creates a new LocationConnectorResponse
func NewLocationConnectorResponse(
	id uuid.UUID,
	connectorID int,
	power float64,
	voltage, amperage int,
	connectorType, connectorStandard string,
) *LocationConnectorResponse {
	return &LocationConnectorResponse{
		ID:                id,
		ConnectorID:       connectorID,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
	}
}
