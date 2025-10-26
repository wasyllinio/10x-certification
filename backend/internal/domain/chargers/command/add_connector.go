package command

import (
	"10x-certification/internal/domain/chargers/dto/request"
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"10x-certification/internal/infrastructure/persistence/models"
	"context"

	"github.com/google/uuid"
)

// AddConnectorCommand represents the command to add a connector to a charger
type AddConnectorCommand struct {
	Request   *request.ConnectorRequest
	ChargerID uuid.UUID
}

// NewAddConnectorCommand creates a new AddConnectorCommand
func NewAddConnectorCommand(chargerID uuid.UUID, req *request.ConnectorRequest) *AddConnectorCommand {
	return &AddConnectorCommand{
		ChargerID: chargerID,
		Request:   req,
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

// HandleInternal - internal method used by CreateChargerHandler for mapping domain to DB model
func (h *AddConnectorHandler) HandleInternal(chargerID uuid.UUID, connector model.Connector) (*models.ConnectorDB, error) {
	// Validate connector
	if err := connector.Validate(); err != nil {
		return nil, err
	}

	// Map domain -> DB
	connectorDB := models.NewConnectorDB()
	connectorDB.ChargerID = chargerID
	connectorDB.ConnectorID = connector.ConnectorID
	connectorDB.Power = float64(connector.Power)
	connectorDB.Voltage = connector.Voltage
	connectorDB.Amperage = connector.Amperage
	connectorDB.ConnectorType = models.ConnectorType(connector.ConnectorType)
	connectorDB.ConnectorStandard = models.ConnectorStandard(connector.ConnectorStandard)

	return connectorDB, nil
}

// Handle executes the add connector command
func (h *AddConnectorHandler) Handle(ctx context.Context, cmd *AddConnectorCommand) error {
	// Create domain model from DTO
	connector := model.NewConnector(
		cmd.ChargerID,
		cmd.Request.ConnectorID,
		float32(cmd.Request.Power),
		cmd.Request.Voltage,
		cmd.Request.Amperage,
		model.ConnectorType(cmd.Request.ConnectorType),
		model.ConnectorStandard(cmd.Request.ConnectorStandard),
	)

	// Map to DB model
	connectorDB, err := h.HandleInternal(cmd.ChargerID, *connector)
	if err != nil {
		return err
	}

	// Save to repository
	return h.connectorRepo.Create(ctx, connectorDB)
}
