package response

import (
	"github.com/google/uuid"
)

// ConnectorResponse represents a connector response DTO
type ConnectorResponse struct {
	ID                uuid.UUID `json:"id"`
	ConnectorID       int       `json:"connector_id"`
	Power             float64   `json:"power"`
	Voltage           int       `json:"voltage"`
	Amperage          int       `json:"amperage"`
	ConnectorType     string    `json:"connector_type"`
	ConnectorStandard string    `json:"connector_standard"`
}

// NewConnectorResponse creates a new ConnectorResponse from domain Connector
func NewConnectorResponse(
	id uuid.UUID,
	connectorID int,
	power float64,
	voltage, amperage int,
	connectorType, connectorStandard string,
) *ConnectorResponse {
	return &ConnectorResponse{
		ID:                id,
		ConnectorID:       connectorID,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
	}
}
