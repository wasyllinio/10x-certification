package query

import (
	"10x-certification/internal/domain/chargers/model"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"
)

// GetLocationChargersQuery represents the query to get chargers for a location
type GetLocationChargersQuery struct {
	LocationID string
}

// NewGetLocationChargersQuery creates a new GetLocationChargersQuery
func NewGetLocationChargersQuery(locationID string) *GetLocationChargersQuery {
	return &GetLocationChargersQuery{
		LocationID: locationID,
	}
}

// GetLocationChargersHandler handles getting chargers for a location
type GetLocationChargersHandler struct {
	locationRepo locationsRepo.LocationRepository
}

// NewGetLocationChargersHandler creates a new GetLocationChargersHandler
func NewGetLocationChargersHandler(locationRepo locationsRepo.LocationRepository) *GetLocationChargersHandler {
	return &GetLocationChargersHandler{
		locationRepo: locationRepo,
	}
}

// Handle executes the get location chargers query
func (h *GetLocationChargersHandler) Handle(ctx context.Context, query *GetLocationChargersQuery) ([]*model.Charger, error) {
	// TODO: Implement get location chargers logic
	panic("not implemented")
}
