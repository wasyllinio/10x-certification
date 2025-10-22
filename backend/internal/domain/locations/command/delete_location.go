package command

import (
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"

	"github.com/google/uuid"
)

// DeleteLocationCommand represents the command to delete a location
type DeleteLocationCommand struct {
	LocationID uuid.UUID
}

// NewDeleteLocationCommand creates a new DeleteLocationCommand
func NewDeleteLocationCommand(locationID uuid.UUID) *DeleteLocationCommand {
	return &DeleteLocationCommand{
		LocationID: locationID,
	}
}

// DeleteLocationHandler handles location deletion
type DeleteLocationHandler struct {
	locationRepo locationsRepo.LocationRepository
}

// NewDeleteLocationHandler creates a new DeleteLocationHandler
func NewDeleteLocationHandler(locationRepo locationsRepo.LocationRepository) *DeleteLocationHandler {
	return &DeleteLocationHandler{
		locationRepo: locationRepo,
	}
}

// Handle executes the delete location command
func (h *DeleteLocationHandler) Handle(ctx context.Context, cmd *DeleteLocationCommand) error {
	// TODO: Implement location deletion logic
	// 1. Check if location has assigned chargers
	// 2. Soft delete location
	// 3. Cascade delete EVSE points
	// 4. Publish LocationDeleted event
	panic("not implemented")
}
