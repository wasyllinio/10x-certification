package query

import (
	"10x-certification/internal/domain/locations/model"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"
)

// ListLocationsQuery represents the query to list locations
type ListLocationsQuery struct {
	OwnerID string
	Page    int
	Limit   int
}

// NewListLocationsQuery creates a new ListLocationsQuery
func NewListLocationsQuery(ownerID string, page, limit int) *ListLocationsQuery {
	return &ListLocationsQuery{
		OwnerID: ownerID,
		Page:    page,
		Limit:   limit,
	}
}

// ListLocationsResult represents the result of listing locations
type ListLocationsResult struct {
	Locations []*model.Location
	Total     int64
	Page      int
	Limit     int
	HasNext   bool
}

// ListLocationsHandler handles listing locations
type ListLocationsHandler struct {
	locationRepo locationsRepo.LocationRepository
}

// NewListLocationsHandler creates a new ListLocationsHandler
func NewListLocationsHandler(locationRepo locationsRepo.LocationRepository) *ListLocationsHandler {
	return &ListLocationsHandler{
		locationRepo: locationRepo,
	}
}

// Handle executes the list locations query
func (h *ListLocationsHandler) Handle(ctx context.Context, query *ListLocationsQuery) (*ListLocationsResult, error) {
	// TODO: Implement list locations logic
	panic("not implemented")
}
