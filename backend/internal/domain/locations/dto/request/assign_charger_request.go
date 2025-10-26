package request

import (
	"github.com/google/uuid"
)

// AssignChargerRequest represents the request to assign a charger to a location
type AssignChargerRequest struct {
	ChargerID uuid.UUID `json:"charger_id" binding:"required"`
}

// NewAssignChargerRequest creates a new AssignChargerRequest
func NewAssignChargerRequest(chargerID uuid.UUID) *AssignChargerRequest {
	return &AssignChargerRequest{
		ChargerID: chargerID,
	}
}
