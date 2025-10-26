package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// ChargerStatus represents charger status enum
type ChargerStatus string

const (
	StatusWarehouse ChargerStatus = "warehouse"
	StatusAssigned  ChargerStatus = "assigned"
)

// Charger represents the charger aggregate root
type Charger struct {
	CreatedAt            time.Time
	UpdatedAt            time.Time
	LocationID           *uuid.UUID
	AssignedToLocationAt *time.Time
	LastStatusChangeAt   *time.Time
	Vendor               string
	Model                string
	SerialNumber         string
	Connectors           []Connector
	Version              int
	ID                   uuid.UUID
	OwnerID              uuid.UUID
}

// NewCharger creates a new charger aggregate
func NewCharger(vendor, model, serialNumber string, ownerID uuid.UUID) *Charger {
	now := time.Now()
	return &Charger{
		Vendor:       vendor,
		Model:        model,
		SerialNumber: serialNumber,
		OwnerID:      ownerID,
		Version:      1,
		Connectors:   make([]Connector, 0),
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// GetStatus returns the current status of the charger
func (c *Charger) GetStatus() ChargerStatus {
	if c.LocationID != nil {
		return StatusAssigned
	}
	return StatusWarehouse
}

// AssignToLocation assigns the charger to a location
func (c *Charger) AssignToLocation(locationID uuid.UUID) error {
	if c.LocationID != nil {
		return errors.New("charger is already assigned to a location")
	}

	now := time.Now()
	c.LocationID = &locationID
	c.AssignedToLocationAt = &now
	c.LastStatusChangeAt = &now
	c.UpdatedAt = now

	return nil
}

// DetachFromLocation detaches the charger from its location
func (c *Charger) DetachFromLocation() error {
	if c.LocationID == nil {
		return errors.New("charger is not assigned to any location")
	}

	now := time.Now()
	c.LocationID = nil
	c.AssignedToLocationAt = nil
	c.LastStatusChangeAt = &now
	c.UpdatedAt = now

	return nil
}

// AddConnector adds a new connector to the charger
func (c *Charger) AddConnector(connector Connector) error {
	// Check if connector ID already exists
	for _, existing := range c.Connectors {
		if existing.ConnectorID == connector.ConnectorID {
			return errors.New("connector ID already exists")
		}
	}

	c.Connectors = append(c.Connectors, connector)
	c.UpdatedAt = time.Now()

	return nil
}

// RemoveConnector removes a connector from the charger
func (c *Charger) RemoveConnector(connectorID int) error {
	for i, connector := range c.Connectors {
		if connector.ConnectorID == connectorID {
			c.Connectors = append(c.Connectors[:i], c.Connectors[i+1:]...)
			c.UpdatedAt = time.Now()
			return nil
		}
	}

	return errors.New("connector not found")
}
