package response

import (
	"time"

	"github.com/google/uuid"
)

// EVSEResponse represents an EVSE response DTO with embedded connector data
type EVSEResponse struct {
	ID        uuid.UUID                 `json:"id"`
	EvseID    string                    `json:"evse_id"`
	Connector LocationConnectorResponse `json:"connector"`
	CreatedAt time.Time                 `json:"created_at"`
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
