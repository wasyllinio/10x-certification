package response

import (
	"10x-certification/internal/infrastructure/http/dto"
)

// LocationListResponse represents a paginated list of locations
type LocationListResponse struct {
	Data       []LocationListItemResponse `json:"data"`
	Pagination dto.PaginationResponse     `json:"pagination"`
}

// NewLocationListResponse creates a new LocationListResponse
func NewLocationListResponse(locations []LocationListItemResponse, pagination dto.PaginationResponse) *LocationListResponse {
	return &LocationListResponse{
		Data:       locations,
		Pagination: pagination,
	}
}
