package request

import "github.com/google/uuid"

// ConnectorRequest represents a connector request DTO
type ConnectorRequest struct {
	ConnectorType     string        `json:"connector_type" binding:"required,oneof=CCS Type2 Chademo"`
	ConnectorStandard string        `json:"connector_standard" binding:"required,oneof=AC_1P AC_3P DC"`
	ConnectorID       int           `json:"connector_id" binding:"required,min=1"`
	Power             float64       `json:"power" binding:"required,gt=0"`
	Voltage           int           `json:"voltage" binding:"required,gt=0"`
	Amperage          int           `json:"amperage" binding:"required,gt=0"`
	ID                uuid.NullUUID `json:"id,omitempty"`
}

// NewConnectorRequest creates a new ConnectorRequest
func NewConnectorRequest(
	id uuid.NullUUID,
	connectorID int,
	power float64,
	voltage, amperage int,
	connectorType, connectorStandard string,
) *ConnectorRequest {
	return &ConnectorRequest{
		ID:                uuid.NullUUID{},
		ConnectorID:       connectorID,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
	}
}
