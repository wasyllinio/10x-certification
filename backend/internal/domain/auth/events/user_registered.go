package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// UserRegistered represents the event when a user is registered
type UserRegistered struct {
	events.BaseEvent
	UserID    uuid.UUID
	Email     string
	Role      string
	Timestamp time.Time
}

// NewUserRegistered creates a new UserRegistered event
func NewUserRegistered(userID uuid.UUID, email, role string) *UserRegistered {
	return &UserRegistered{
		BaseEvent: events.NewBaseEvent("UserRegistered"),
		UserID:    userID,
		Email:     email,
		Role:      role,
		Timestamp: time.Now(),
	}
}
