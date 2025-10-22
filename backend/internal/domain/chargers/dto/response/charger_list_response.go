package response

import (
	"10x-certification/internal/infrastructure/http/dto"
)

// ChargerListResponse represents a paginated list of chargers
type ChargerListResponse struct {
	Data       []ChargerListItemResponse `json:"data"`
	Pagination dto.PaginationResponse    `json:"pagination"`
}

// NewChargerListResponse creates a new ChargerListResponse
func NewChargerListResponse(chargers []ChargerListItemResponse, pagination dto.PaginationResponse) *ChargerListResponse {
	return &ChargerListResponse{
		Data:       chargers,
		Pagination: pagination,
	}
}
