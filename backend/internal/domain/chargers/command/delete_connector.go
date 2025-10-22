package command

import (
	"10x-certification/internal/domain/chargers/repository"
	"context"

	"github.com/google/uuid"
)

// DeleteConnectorCommand represents the command to delete a connector
type DeleteConnectorCommand struct {
	ConnectorID uuid.UUID
}

// NewDeleteConnectorCommand creates a new DeleteConnectorCommand
func NewDeleteConnectorCommand(connectorID uuid.UUID) *DeleteConnectorCommand {
	return &DeleteConnectorCommand{
		ConnectorID: connectorID,
	}
}

// DeleteConnectorHandler handles connector deletion
type DeleteConnectorHandler struct {
	connectorRepo repository.ConnectorRepository
}

// NewDeleteConnectorHandler creates a new DeleteConnectorHandler
func NewDeleteConnectorHandler(connectorRepo repository.ConnectorRepository) *DeleteConnectorHandler {
	return &DeleteConnectorHandler{
		connectorRepo: connectorRepo,
	}
}

// Handle executes the delete connector command
func (h *DeleteConnectorHandler) Handle(ctx context.Context, cmd *DeleteConnectorCommand) error {
	// TODO: Implement connector deletion logic
	// 1. Check if connector has associated EVSE
	// 2. Soft delete connector
	// 3. Publish ConnectorRemoved event
	panic("not implemented")
}
