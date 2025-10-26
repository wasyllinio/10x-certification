package query

import (
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"context"
)

// SearchChargersQuery represents the query to search chargers
type SearchChargersQuery struct {
	OwnerID string
	Search  string
	Status  string
	Page    int
	Limit   int
}

// NewSearchChargersQuery creates a new SearchChargersQuery
func NewSearchChargersQuery(ownerID, search string, page, limit int, status string) *SearchChargersQuery {
	return &SearchChargersQuery{
		OwnerID: ownerID,
		Search:  search,
		Page:    page,
		Limit:   limit,
		Status:  status,
	}
}

// SearchChargersResult represents the result of searching chargers
type SearchChargersResult struct {
	Chargers []*model.Charger
	Total    int64
	Page     int
	Limit    int
	HasNext  bool
}

// SearchChargersHandler handles searching chargers
type SearchChargersHandler struct {
	chargerRepo repository.ChargerRepository
}

// NewSearchChargersHandler creates a new SearchChargersHandler
func NewSearchChargersHandler(chargerRepo repository.ChargerRepository) *SearchChargersHandler {
	return &SearchChargersHandler{
		chargerRepo: chargerRepo,
	}
}

// Handle executes the search chargers query
func (h *SearchChargersHandler) Handle(ctx context.Context, query *SearchChargersQuery) (*SearchChargersResult, error) {
	// TODO: Implement search chargers logic
	panic("not implemented")
}
