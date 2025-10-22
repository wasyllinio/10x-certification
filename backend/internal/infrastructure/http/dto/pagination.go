package dto

import (
	"strconv"
)

// PaginationRequest represents pagination request parameters
type PaginationRequest struct {
	Page  int `form:"page" binding:"min=1"`
	Limit int `form:"limit" binding:"min=1,max=100"`
}

// PaginationResponse represents pagination response
type PaginationResponse struct {
	Page    int   `json:"page"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	HasNext bool  `json:"has_next"`
}

// NewPaginationRequest creates a new pagination request from query parameters
func NewPaginationRequest(pageStr, limitStr string) PaginationRequest {
	page := 1
	limit := 20

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
		limit = l
	}

	return PaginationRequest{
		Page:  page,
		Limit: limit,
	}
}

// ToResponse converts pagination request to response
func (p *PaginationRequest) ToResponse(total int64) PaginationResponse {
	return PaginationResponse{
		Page:    p.Page,
		Limit:   p.Limit,
		Total:   total,
		HasNext: int64(p.Page*p.Limit) < total,
	}
}
