package events

import (
	"time"
)

// BaseEvent represents the base event interface
type BaseEvent interface {
	GetEventType() string
	GetTimestamp() time.Time
	GetEventID() string
}

// baseEvent represents the base implementation of an event
type baseEvent struct {
	eventType string
	timestamp time.Time
	eventID   string
}

// NewBaseEvent creates a new base event
func NewBaseEvent(eventType string) BaseEvent {
	return &baseEvent{
		eventType: eventType,
		timestamp: time.Now(),
		eventID:   "generated-event-id", // TODO: Generate actual event ID
	}
}

// GetEventType returns the event type
func (e *baseEvent) GetEventType() string {
	return e.eventType
}

// GetTimestamp returns the event timestamp
func (e *baseEvent) GetTimestamp() time.Time {
	return e.timestamp
}

// GetEventID returns the event ID
func (e *baseEvent) GetEventID() string {
	return e.eventID
}
