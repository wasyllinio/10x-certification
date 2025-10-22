package query

import (
	"10x-certification/internal/domain/chargers/model"
	"10x-certification/internal/domain/chargers/repository"
	"context"
)

// GetChargerByIDQuery represents the query to get charger by ID
type GetChargerByIDQuery struct {
	ChargerID string
}

// NewGetChargerByIDQuery creates a new GetChargerByIDQuery
func NewGetChargerByIDQuery(chargerID string) *GetChargerByIDQuery {
	return &GetChargerByIDQuery{
		ChargerID: chargerID,
	}
}

// GetChargerByIDHandler handles getting charger by ID
type GetChargerByIDHandler struct {
	chargerRepo repository.ChargerRepository
}

// NewGetChargerByIDHandler creates a new GetChargerByIDHandler
func NewGetChargerByIDHandler(chargerRepo repository.ChargerRepository) *GetChargerByIDHandler {
	return &GetChargerByIDHandler{
		chargerRepo: chargerRepo,
	}
}

// Handle executes the get charger by ID query
func (h *GetChargerByIDHandler) Handle(ctx context.Context, query *GetChargerByIDQuery) (*model.Charger, error) {
	// TODO: Implement get charger by ID logic
	panic("not implemented")
}
