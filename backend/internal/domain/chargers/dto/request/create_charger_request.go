package request

// CreateChargerRequest represents the request to create a new charger
type CreateChargerRequest struct {
	Vendor       string             `json:"vendor" validate:"required"`
	Model        string             `json:"model" validate:"required"`
	SerialNumber string             `json:"serial_number" validate:"required"`
	Connectors   []ConnectorRequest `json:"connectors" validate:"required,min=1,dive"`
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
