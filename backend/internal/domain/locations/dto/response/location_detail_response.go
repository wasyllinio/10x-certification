package response

import "time"

// LocationDetailResponse represents a detailed location response DTO with embedded chargers and EVSEs
type LocationDetailResponse struct {
	Name        string                    `json:"name"`
	Address     string                    `json:"address"`
	CountryCode string                    `json:"country_code"`
	Chargers    []LocationChargerResponse `json:"chargers"`
	EVSEs       []EVSEResponse            `json:"evses"`
	Version     int                       `json:"version"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
}

// NewLocationDetailResponse creates a new LocationDetailResponse from domain Location
func NewLocationDetailResponse(
	name, address, countryCode string,
	version int,
	chargers []LocationChargerResponse,
	evses []EVSEResponse,
	createdAt, updatedAt time.Time,
) *LocationDetailResponse {
	return &LocationDetailResponse{
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		Version:     version,
		Chargers:    chargers,
		EVSEs:       evses,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
