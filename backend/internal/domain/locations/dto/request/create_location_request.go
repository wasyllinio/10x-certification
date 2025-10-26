package request

// CreateLocationRequest represents the request to create a new location
type CreateLocationRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	CountryCode string `json:"country_code" binding:"required,len=3,regexp=^[A-Z]{3}$"`
}

// NewCreateLocationRequest creates a new CreateLocationRequest
func NewCreateLocationRequest(name, address, countryCode string) *CreateLocationRequest {
	return &CreateLocationRequest{
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
	}
}
