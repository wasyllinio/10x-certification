package request

import (
	"github.com/google/uuid"
)

// SearchChargersRequest represents search and filter parameters for chargers
type SearchChargersRequest struct {
	LocationID *uuid.UUID `json:"location_id,omitempty"`
	Search     string     `json:"search,omitempty"`
	Status     string     `json:"status,omitempty" binding:"omitempty,oneof=warehouse assigned"`
	Page       int        `json:"page" binding:"min=1"`
	Limit      int        `json:"limit" binding:"min=1,max=100"`
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
