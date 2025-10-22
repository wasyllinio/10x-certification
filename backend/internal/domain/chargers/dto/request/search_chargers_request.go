package request

import (
	"github.com/google/uuid"
)

// SearchChargersRequest represents search and filter parameters for chargers
type SearchChargersRequest struct {
	Page       int        `json:"page" validate:"min=1"`
	Limit      int        `json:"limit" validate:"min=1,max=100"`
	Search     string     `json:"search,omitempty"`
	Status     string     `json:"status,omitempty" validate:"omitempty,oneof=warehouse assigned"`
	LocationID *uuid.UUID `json:"location_id,omitempty"`
}

// NewSearchChargersRequest creates a new SearchChargersRequest
func NewSearchChargersRequest(page, limit int, search, status string, locationID *uuid.UUID) *SearchChargersRequest {
	return &SearchChargersRequest{
		Page:       page,
		Limit:      limit,
		Search:     search,
		Status:     status,
		LocationID: locationID,
	}
}

// GetPage returns the page number, defaulting to 1
func (s *SearchChargersRequest) GetPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}

// GetLimit returns the limit, defaulting to 20
func (s *SearchChargersRequest) GetLimit() int {
	if s.Limit <= 0 {
		return 20
	}
	return s.Limit
}
