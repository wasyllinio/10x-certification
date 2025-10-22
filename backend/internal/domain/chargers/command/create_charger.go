package command

import (
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"context"

	"github.com/google/uuid"
)

// CreateChargerCommand represents the command to create a new charger
type CreateChargerCommand struct {
	Vendor       string
	Model        string
	SerialNumber string
	OwnerID      uuid.UUID
	Connectors   []model.Connector
}

// NewCreateChargerCommand creates a new CreateChargerCommand
func NewCreateChargerCommand(vendor, model, serialNumber string, ownerID uuid.UUID, connectors []model.Connector) *CreateChargerCommand {
	return &CreateChargerCommand{
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		Connectors:   connectors,
	}
}

// CreateChargerHandler handles charger creation
type CreateChargerHandler struct {
	chargerRepo repository.ChargerRepository
}

// NewCreateChargerHandler creates a new CreateChargerHandler
func NewCreateChargerHandler(chargerRepo repository.ChargerRepository, connectorRepo repository.ConnectorRepository) *CreateChargerHandler {
	return &CreateChargerHandler{
		chargerRepo: chargerRepo,
	}
}

// Handle executes the create charger command
func (h *CreateChargerHandler) Handle(ctx context.Context, cmd *CreateChargerCommand) (*model.Charger, error) {
	// TODO: Implement charger creation logic
	// 1. Validate charger data
	// 2. Create charger aggregate
	// 3. Save to repository
	// 4. Publish ChargerCreated event
	panic("not implemented")
}
