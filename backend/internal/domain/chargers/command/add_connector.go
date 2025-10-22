package command

import (
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"context"

	"github.com/google/uuid"
)

// AddConnectorCommand represents the command to add a connector to a charger
type AddConnectorCommand struct {
	ChargerID         uuid.UUID
	ConnectorID       int
	Power             float32
	Voltage           int
	Amperage          int
	ConnectorType     model.ConnectorType
	ConnectorStandard model.ConnectorStandard
}

// NewAddConnectorCommand creates a new AddConnectorCommand
func NewAddConnectorCommand(
	chargerID uuid.UUID,
	connectorID int,
	power float32,
	voltage, amperage int,
	connectorType model.ConnectorType,
	connectorStandard model.ConnectorStandard,
) *AddConnectorCommand {
	return &AddConnectorCommand{
		ChargerID:         chargerID,
		ConnectorID:       connectorID,
		Power:             power,
		Voltage:           voltage,
		Amperage:          amperage,
		ConnectorType:     connectorType,
		ConnectorStandard: connectorStandard,
	}
}

// AddConnectorHandler handles adding connectors to chargers
type AddConnectorHandler struct {
	connectorRepo repository.ConnectorRepository
}

// NewAddConnectorHandler creates a new AddConnectorHandler
func NewAddConnectorHandler(connectorRepo repository.ConnectorRepository) *AddConnectorHandler {
	return &AddConnectorHandler{
		connectorRepo: connectorRepo,
	}
}

// Handle executes the add connector command
func (h *AddConnectorHandler) Handle(ctx context.Context, cmd *AddConnectorCommand) (*model.Charger, error) {
	// TODO: Implement add connector logic
	// 1. Find charger by ID
	// 2. Create new connector
	// 3. Add connector to charger
	// 4. Save to repository
	// 5. Publish ConnectorAdded event
	panic("not implemented")
}
