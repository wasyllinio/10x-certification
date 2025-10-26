package response

import (
	"time"

	"github.com/google/uuid"
)

// EVSEResponse represents an EVSE response DTO with embedded connector data
type EVSEResponse struct {
	CreatedAt time.Time                 `json:"created_at"`
	EvseID    string                    `json:"evse_id"`
	Connector LocationConnectorResponse `json:"connector"`
	ID        uuid.UUID                 `json:"id"`
}

// NewEVSEResponse creates a new EVSEResponse from domain EVSE
func NewEVSEResponse(
	id uuid.UUID,
	evseID string,
	connector LocationConnectorResponse,
	createdAt time.Time,
) *EVSEResponse {
	return &EVSEResponse{
		ID:        id,
		EvseID:    evseID,
		Connector: connector,
		CreatedAt: createdAt,
	}
}
