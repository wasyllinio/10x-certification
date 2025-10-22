package request

// UpdateLocationRequest represents the request to update a location
type UpdateLocationRequest struct {
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	CountryCode string `json:"country_code" validate:"required,len=3,regexp=^[A-Z]{3}$"`
	Version     int    `json:"version" validate:"required,min=1"`
}

// NewUpdateLocationRequest creates a new UpdateLocationRequest
func NewUpdateLocationRequest(name, address, countryCode string, version int) *UpdateLocationRequest {
	return &UpdateLocationRequest{
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		Version:     version,
	}
}
