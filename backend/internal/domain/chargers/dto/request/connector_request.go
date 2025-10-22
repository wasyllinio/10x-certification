package request

// ConnectorRequest represents a connector request DTO
type ConnectorRequest struct {
	ConnectorID       int     `json:"connector_id" validate:"required,min=1"`
	Power             float64 `json:"power" validate:"required,gt=0"`
	Voltage           int     `json:"voltage" validate:"required,gt=0"`
	Amperage          int     `json:"amperage" validate:"required,gt=0"`
	ConnectorType     string  `json:"connector_type" validate:"required,oneof=CCS Type2 Chademo"`
	ConnectorStandard string  `json:"connector_standard" validate:"required,oneof=AC_1P AC_3P DC"`
}

// NewConnectorRequest creates a new ConnectorRequest
func NewConnectorRequest(
	connectorID int,
	power float64,
	voltage, amperage int,
	connectorType, connectorStandard string,
) *ConnectorRequest {
	return &ConnectorRequest{
		ConnectorID:       connectorID,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
	}
}
