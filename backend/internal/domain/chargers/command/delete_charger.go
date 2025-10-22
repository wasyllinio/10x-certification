package command

import (
	"10x-certification/internal/domain/chargers/repository"
	"context"

	"github.com/google/uuid"
)

// DeleteChargerCommand represents the command to delete a charger
type DeleteChargerCommand struct {
	ChargerID uuid.UUID
}

// NewDeleteChargerCommand creates a new DeleteChargerCommand
func NewDeleteChargerCommand(chargerID uuid.UUID) *DeleteChargerCommand {
	return &DeleteChargerCommand{
		ChargerID: chargerID,
	}
}

// DeleteChargerHandler handles charger deletion
type DeleteChargerHandler struct {
	chargerRepo repository.ChargerRepository
}

// NewDeleteChargerHandler creates a new DeleteChargerHandler
func NewDeleteChargerHandler(chargerRepo repository.ChargerRepository) *DeleteChargerHandler {
	return &DeleteChargerHandler{
		chargerRepo: chargerRepo,
	}
}

// Handle executes the delete charger command
func (h *DeleteChargerHandler) Handle(ctx context.Context, cmd *DeleteChargerCommand) error {
	// TODO: Implement charger deletion logic
	// 1. Find charger by ID
	// 2. Check if charger is assigned to location (cannot delete if assigned)
	// 3. Soft delete charger
	// 4. Save to repository
	// 5. Publish ChargerDeleted event
	panic("not implemented")
}
