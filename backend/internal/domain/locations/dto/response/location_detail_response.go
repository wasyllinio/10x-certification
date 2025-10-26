package response

import (
	"github.com/google/uuid"
)

// LocationDetailResponse represents a detailed location response DTO with embedded chargers and EVSEs
type LocationDetailResponse struct {
	Name        string                    `json:"name"`
	Address     string                    `json:"address"`
	CountryCode string                    `json:"country_code"`
	Chargers    []LocationChargerResponse `json:"chargers"`
	EVSEs       []EVSEResponse            `json:"evses"`
	Version     int                       `json:"version"`
	ID          uuid.UUID                 `json:"id"`
}

// NewLocationDetailResponse creates a new LocationDetailResponse from domain Location
func NewLocationDetailResponse(
	id uuid.UUID,
	name, address, countryCode string,
	version int,
	chargers []LocationChargerResponse,
	evses []EVSEResponse,
) *LocationDetailResponse {
	return &LocationDetailResponse{
		ID:          id,
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		Version:     version,
		Chargers:    chargers,
		EVSEs:       evses,
	}
}
