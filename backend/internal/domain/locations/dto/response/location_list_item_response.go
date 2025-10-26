package response

import (
	"time"

	"github.com/google/uuid"
)

// LocationListItemResponse represents a location item in a list response
type LocationListItemResponse struct {
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	CountryCode   string    `json:"country_code"`
	ChargersCount int       `json:"chargers_count"`
	EVSECount     int       `json:"evse_count"`
	ID            uuid.UUID `json:"id"`
	OwnerID       uuid.UUID `json:"owner_id"`
}

// NewLocationListItemResponse creates a new LocationListItemResponse from domain Location
func NewLocationListItemResponse(
	id, ownerID uuid.UUID,
	name, address, countryCode string,
	chargersCount, evseCount int,
	createdAt, updatedAt time.Time,
) *LocationListItemResponse {
	return &LocationListItemResponse{
		ID:            id,
		Name:          name,
		Address:       address,
		CountryCode:   countryCode,
		OwnerID:       ownerID,
		ChargersCount: chargersCount,
		EVSECount:     evseCount,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}
