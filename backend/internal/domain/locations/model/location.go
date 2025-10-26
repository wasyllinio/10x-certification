package model

import (
	"time"

	"github.com/google/uuid"
)

// Location represents the location aggregate root
type Location struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Address     string
	CountryCode string
	Version     int
	ID          uuid.UUID
	OwnerID     uuid.UUID
}

// NewLocation creates a new location aggregate
func NewLocation(name, address string, countryCode string, ownerID uuid.UUID) *Location {
	now := time.Now()
	return &Location{
		ID:          uuid.New(),
		Name:        name,
		Address:     address,
		CountryCode: countryCode,
		OwnerID:     ownerID,
		Version:     1,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// Update updates location data
func (l *Location) Update(name, address string, countryCode string) {
	l.Name = name
	l.Address = address
	l.CountryCode = countryCode
	l.Version++
	l.UpdatedAt = time.Now()
}
