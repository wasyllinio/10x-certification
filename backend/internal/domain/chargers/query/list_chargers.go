package query

import (
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"context"
)

// ListChargersQuery represents the query to list chargers
type ListChargersQuery struct {
	OwnerID string
	Page    int
	Limit   int
	Status  string
}

// NewListChargersQuery creates a new ListChargersQuery
func NewListChargersQuery(ownerID string, page, limit int, status string) *ListChargersQuery {
	return &ListChargersQuery{
		OwnerID: ownerID,
		Page:    page,
		Limit:   limit,
		Status:  status,
	}
}

// ListChargersResult represents the result of listing chargers
type ListChargersResult struct {
	Chargers []*model.Charger
	Total    int64
	Page     int
	Limit    int
	HasNext  bool
}

// ListChargersHandler handles listing chargers
type ListChargersHandler struct {
	chargerRepo repository.ChargerRepository
}

// NewListChargersHandler creates a new ListChargersHandler
func NewListChargersHandler(chargerRepo repository.ChargerRepository) *ListChargersHandler {
	return &ListChargersHandler{
		chargerRepo: chargerRepo,
	}
}

// Handle executes the list chargers query
func (h *ListChargersHandler) Handle(ctx context.Context, query *ListChargersQuery) (*ListChargersResult, error) {
	// TODO: Implement list chargers logic
	panic("not implemented")
}
