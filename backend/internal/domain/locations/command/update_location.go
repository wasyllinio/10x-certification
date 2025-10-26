package command

import (
	"10x-certification/internal/domain/locations/dto/request"
	"10x-certification/internal/domain/locations/model"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"

	"github.com/google/uuid"
)

// UpdateLocationCommand represents the command to update a location
type UpdateLocationCommand struct {
	Request    *request.UpdateLocationRequest
	LocationID uuid.UUID
}

// NewUpdateLocationCommand creates a new UpdateLocationCommand
func NewUpdateLocationCommand(locationID uuid.UUID, req *request.UpdateLocationRequest) *UpdateLocationCommand {
	return &UpdateLocationCommand{
		LocationID: locationID,
		Request:    req,
	}
}

// UpdateLocationHandler handles location updates
type UpdateLocationHandler struct {
	locationRepo locationsRepo.LocationRepository
}

// NewUpdateLocationHandler creates a new UpdateLocationHandler
func NewUpdateLocationHandler(locationRepo locationsRepo.LocationRepository) *UpdateLocationHandler {
	return &UpdateLocationHandler{
		locationRepo: locationRepo,
	}
}

// Handle executes the update location command
func (h *UpdateLocationHandler) Handle(ctx context.Context, cmd *UpdateLocationCommand) (*model.Location, error) {
	// TODO: Implement location update logic
	// 1. Get existing location
	// 2. Validate optimistic locking
	// 3. Update location data
	// 4. Save to repository
	// 5. Publish LocationUpdated event
	panic("not implemented")
}
