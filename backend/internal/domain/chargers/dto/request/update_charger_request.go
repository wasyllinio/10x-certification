package request

// UpdateChargerRequest represents the request to update a charger
type UpdateChargerRequest struct {
	Vendor       string             `json:"vendor" validate:"required"`
	Model        string             `json:"model" validate:"required"`
	SerialNumber string             `json:"serial_number" validate:"required"`
	Version      int                `json:"version" validate:"required,min=1"`
	Connectors   []ConnectorRequest `json:"connectors" validate:"required,min=1,dive"`
}

// NewUpdateChargerRequest creates a new UpdateChargerRequest
func NewUpdateChargerRequest(vendor, model, serialNumber string, version int, connectors []ConnectorRequest) *UpdateChargerRequest {
	return &UpdateChargerRequest{
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		Version:      version,
		Connectors:   connectors,
	}
}
