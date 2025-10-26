package request

// CreateChargerRequest represents the request to create a new charger
type CreateChargerRequest struct {
	Vendor       string             `json:"vendor" binding:"required"`
	Model        string             `json:"model" binding:"required"`
	SerialNumber string             `json:"serial_number" binding:"required"`
	Connectors   []ConnectorRequest `json:"connectors" binding:"required,min=1,dive"`
}

// NewCreateChargerRequest creates a new CreateChargerRequest
func NewCreateChargerRequest(vendor, model, serialNumber string, connectors []ConnectorRequest) *CreateChargerRequest {
	return &CreateChargerRequest{
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		Connectors:   connectors,
	}
}
