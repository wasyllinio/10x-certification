package query

import (
	"10x-certification/internal/domain/locations/model"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"
)

// GetLocationByIDQuery represents the query to get location by ID
type GetLocationByIDQuery struct {
	LocationID string
}

// NewGetLocationByIDQuery creates a new GetLocationByIDQuery
func NewGetLocationByIDQuery(locationID string) *GetLocationByIDQuery {
	return &GetLocationByIDQuery{
		LocationID: locationID,
	}
}

// GetLocationByIDHandler handles getting location by ID
type GetLocationByIDHandler struct {
	locationRepo locationsRepo.LocationRepository
}

// NewGetLocationByIDHandler creates a new GetLocationByIDHandler
func NewGetLocationByIDHandler(locationRepo locationsRepo.LocationRepository) *GetLocationByIDHandler {
	return &GetLocationByIDHandler{
		locationRepo: locationRepo,
	}
}

// Handle executes the get location by ID query
func (h *GetLocationByIDHandler) Handle(ctx context.Context, query *GetLocationByIDQuery) (*model.Location, error) {
	// TODO: Implement get location by ID logic
	panic("not implemented")
}
