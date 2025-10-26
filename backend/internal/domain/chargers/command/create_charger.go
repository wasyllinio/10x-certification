package command

import (
	"10x-certification/internal/domain/chargers/dto/request"
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"10x-certification/internal/infrastructure/persistence/models"
	"10x-certification/internal/shared/errors"
	"context"

	"github.com/google/uuid"
)

// CreateChargerCommand represents the command to create a new charger
type CreateChargerCommand struct {
	Request *request.CreateChargerRequest
	OwnerID uuid.UUID
}

// NewCreateChargerCommand creates a new CreateChargerCommand
func NewCreateChargerCommand(req *request.CreateChargerRequest, ownerID uuid.UUID) *CreateChargerCommand {
	return &CreateChargerCommand{
		Request: req,
		OwnerID: ownerID,
	}
}

// CreateChargerHandler handles charger creation
type CreateChargerHandler struct {
	chargerRepo         repository.ChargerRepository
	addConnectorHandler *AddConnectorHandler
}

// NewCreateChargerHandler creates a new CreateChargerHandler
func NewCreateChargerHandler(chargerRepo repository.ChargerRepository, addConnectorHandler *AddConnectorHandler) *CreateChargerHandler {
	return &CreateChargerHandler{
		chargerRepo:         chargerRepo,
		addConnectorHandler: addConnectorHandler,
	}
}

// Handle executes the create charger command
func (h *CreateChargerHandler) Handle(ctx context.Context, cmd *CreateChargerCommand) (*uuid.UUID, error) {
	// 1. Check for duplicates
	exists, err := h.chargerRepo.ExistsByVendorAndSerial(ctx, cmd.Request.Vendor, cmd.Request.SerialNumber)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.ErrChargerAlreadyExists
	}

	// 2. Create charger aggregate
	charger := model.NewCharger(cmd.Request.Vendor, cmd.Request.Model, cmd.Request.SerialNumber, cmd.OwnerID)

	// 3. Add connectors to charger aggregate
	for _, connectorReq := range cmd.Request.Connectors {
		connector := model.NewConnector(
			charger.ID, // Now we have charger ID
			connectorReq.ConnectorID,
			float32(connectorReq.Power),
			connectorReq.Voltage,
			connectorReq.Amperage,
			model.ConnectorType(connectorReq.ConnectorType),
			model.ConnectorStandard(connectorReq.ConnectorStandard),
		)
		if err := charger.AddConnector(*connector); err != nil {
			return nil, err
		}
	}

	// 4. Prepare DB model charger
	chargerDB := models.NewChargerDB()
	chargerDB.ID = charger.ID
	chargerDB.Vendor = charger.Vendor
	chargerDB.Model = charger.Model
	chargerDB.SerialNumber = charger.SerialNumber
	chargerDB.OwnerID = charger.OwnerID
	chargerDB.Version = charger.Version

	// 5. Use AddConnectorHandler to prepare connectors
	connectorsDB := make([]models.ConnectorDB, len(charger.Connectors))
	for i, connector := range charger.Connectors {
		connectorDB, err := h.addConnectorHandler.HandleInternal(charger.ID, connector)
		if err != nil {
			return nil, err
		}
		connectorsDB[i] = *connectorDB
	}

	// 6. Save in transaction
	if err := h.chargerRepo.Create(ctx, chargerDB, connectorsDB); err != nil {
		return nil, err
	}

	return &charger.ID, nil
}
