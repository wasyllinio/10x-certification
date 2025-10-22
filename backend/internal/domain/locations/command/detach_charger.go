package command

import (
	chargersRepo "10x-certification/internal/domain/chargers/repository"
	locationsRepo "10x-certification/internal/domain/locations/repository"
	"context"

	"github.com/google/uuid"
)

// DetachChargerCommand represents the command to detach charger from location
type DetachChargerCommand struct {
	LocationID uuid.UUID
	ChargerID  uuid.UUID
}

// NewDetachChargerCommand creates a new DetachChargerCommand
func NewDetachChargerCommand(locationID, chargerID uuid.UUID) *DetachChargerCommand {
	return &DetachChargerCommand{
		LocationID: locationID,
		ChargerID:  chargerID,
	}
}

// DetachChargerHandler handles charger detachment from location
type DetachChargerHandler struct {
	locationRepo locationsRepo.LocationRepository
	chargerRepo  chargersRepo.ChargerRepository
}

// NewDetachChargerHandler creates a new DetachChargerHandler
func NewDetachChargerHandler(locationRepo locationsRepo.LocationRepository, chargerRepo chargersRepo.ChargerRepository) *DetachChargerHandler {
	return &DetachChargerHandler{
		locationRepo: locationRepo,
		chargerRepo:  chargerRepo,
	}
}

// Handle executes the detach charger command
func (h *DetachChargerHandler) Handle(ctx context.Context, cmd *DetachChargerCommand) error {
	// TODO: Implement charger detachment logic
	// 1. Get location and charger
	// 2. Validate detachment rules
	// 3. Detach charger from location
	// 4. Delete associated EVSE points
	// 5. Save changes
	// 6. Publish ChargerDetached event
	panic("not implemented")
}
