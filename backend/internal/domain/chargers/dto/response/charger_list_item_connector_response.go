package response

// ChargerListItemConnectorResponse represents a connector response DTO for charger list items
type ChargerListItemConnectorResponse struct {
	ConnectorType string  `json:"connector_type"`
	Power         float64 `json:"power"`
}

// NewChargerListItemConnectorResponse creates a new ChargerListItemConnectorResponse
func NewChargerListItemConnectorResponse(
	power float64,
	connectorType string,
) *ChargerListItemConnectorResponse {
	return &ChargerListItemConnectorResponse{
		Power:         power,
		ConnectorType: connectorType,
	}
}
