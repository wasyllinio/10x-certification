package command

import (
	"10x-certification/internal/domain/locations/model"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"

	"github.com/google/uuid"
)

// CreateLocationCommand represents the command to create a new location
type CreateLocationCommand struct {
	Name        string
	Address     string
	CountryCode string
	OwnerID     uuid.UUID
}

// NewCreateLocationCommand creates a new CreateLocationCommand
func NewCreateLocationCommand(name, address, countryCode string, ownerID uuid.UUID) *CreateLocationCommand {
	return &CreateLocationCommand{
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		OwnerID:     ownerID,
	}
}

// CreateLocationHandler handles location creation
type CreateLocationHandler struct {
	locationRepo locationsRepo.LocationRepository
}

// NewCreateLocationHandler creates a new CreateLocationHandler
func NewCreateLocationHandler(locationRepo locationsRepo.LocationRepository) *CreateLocationHandler {
	return &CreateLocationHandler{
		locationRepo: locationRepo,
	}
}

// Handle executes the create location command
func (h *CreateLocationHandler) Handle(ctx context.Context, cmd *CreateLocationCommand) (*model.Location, error) {
	// TODO: Implement location creation logic
	// 1. Validate location data
	// 2. Create location aggregate
	// 3. Save to repository
	// 4. Publish LocationCreated event
	panic("not implemented")
}
