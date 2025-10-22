package model

import (
	"time"

	"github.com/google/uuid"
)

// Location represents the location aggregate root
type Location struct {
	ID          uuid.UUID
	Name        string
	Address     string
	CountryCode string
	OwnerID     uuid.UUID
	Version     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
