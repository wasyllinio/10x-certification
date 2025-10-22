package service

import (
	"10x-certification/internal/domain/locations/model"
	"10x-certification/internal/domain/locations/repository"
)

// LocationService represents domain service for locations
type LocationService struct {
	locationRepo repository.LocationRepository
}

// NewLocationService creates a new LocationService
func NewLocationService(locationRepo repository.LocationRepository) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

// ValidateLocationDeletion validates if location can be deleted
func (s *LocationService) ValidateLocationDeletion(locationID string) error {
	// TODO: Implement location deletion validation logic
	// Check if location has assigned chargers
	panic("not implemented")
}

// EVSEGeneratorService represents domain service for EVSE generation
type EVSEGeneratorService struct {
	evseRepo repository.EVSERepository
}

// NewEVSEGeneratorService creates a new EVSEGeneratorService
func NewEVSEGeneratorService(evseRepo repository.EVSERepository) *EVSEGeneratorService {
	return &EVSEGeneratorService{
		evseRepo: evseRepo,
	}
}

// GenerateEVSEFromConnector generates EVSE from connector
func (s *EVSEGeneratorService) GenerateEVSEFromConnector(connectorID, locationID string) (*model.EVSE, error) {
	// TODO: Implement EVSE generation logic
	// 1. Generate EvseID according to Emi3spec
	// 2. Create EVSE entity
	// 3. Save to repository
	panic("not implemented")
}
