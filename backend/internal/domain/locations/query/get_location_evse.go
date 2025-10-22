package query

import (
	"10x-certification/internal/domain/locations/model"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"
)

// GetLocationEVSEQuery represents the query to get EVSE for a location
type GetLocationEVSEQuery struct {
	LocationID string
}

// NewGetLocationEVSEQuery creates a new GetLocationEVSEQuery
func NewGetLocationEVSEQuery(locationID string) *GetLocationEVSEQuery {
	return &GetLocationEVSEQuery{
		LocationID: locationID,
	}
}

// GetLocationEVSEHandler handles getting EVSE for a location
type GetLocationEVSEHandler struct {
	evseRepo locationsRepo.EVSERepository
}

// NewGetLocationEVSEHandler creates a new GetLocationEVSEHandler
func NewGetLocationEVSEHandler(evseRepo locationsRepo.EVSERepository) *GetLocationEVSEHandler {
	return &GetLocationEVSEHandler{
		evseRepo: evseRepo,
	}
}

// Handle executes the get location EVSE query
func (h *GetLocationEVSEHandler) Handle(ctx context.Context, query *GetLocationEVSEQuery) ([]*model.EVSE, error) {
	// TODO: Implement get location EVSE logic
	panic("not implemented")
}
