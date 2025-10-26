package request

// UpdateChargerRequest represents the request to update a charger
type UpdateChargerRequest struct {
	Vendor       string             `json:"vendor" binding:"required"`
	Model        string             `json:"model" binding:"required"`
	SerialNumber string             `json:"serial_number" binding:"required"`
	Connectors   []ConnectorRequest `json:"connectors" binding:"required,min=1,dive"`
	Version      int                `json:"version" binding:"required,min=1"`
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
