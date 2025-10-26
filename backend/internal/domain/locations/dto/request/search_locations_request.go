package request

// SearchLocationsRequest represents search and filter parameters for locations
type SearchLocationsRequest struct {
	Search      string `json:"search,omitempty"`
	CountryCode string `json:"country_code,omitempty" binding:"omitempty,len=3,regexp=^[A-Z]{3}$"`
	Page        int    `json:"page" binding:"min=1"`
	Limit       int    `json:"limit" binding:"min=1,max=100"`
}

// NewSearchLocationsRequest creates a new SearchLocationsRequest
func NewSearchLocationsRequest(page, limit int, search, countryCode string) *SearchLocationsRequest {
	return &SearchLocationsRequest{
		Page:        page,
		Limit:       limit,
		Search:      search,
		CountryCode: countryCode,
	}
}

// GetPage returns the page number, defaulting to 1
func (s *SearchLocationsRequest) GetPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}

// GetLimit returns the limit, defaulting to 20
func (s *SearchLocationsRequest) GetLimit() int {
	if s.Limit <= 0 {
		return 20
	}
	return s.Limit
}
