package events

import (
	"10x-certification/internal/shared/events"
	"time"

	"github.com/google/uuid"
)

// UserLoggedIn represents the event when a user logs in
type UserLoggedIn struct {
	Timestamp time.Time
	events.BaseEvent
	Email  string
	UserID uuid.UUID
}

// NewUserLoggedIn creates a new UserLoggedIn event
func NewUserLoggedIn(userID uuid.UUID, email string) *UserLoggedIn {
	return &UserLoggedIn{
		BaseEvent: events.NewBaseEvent("UserLoggedIn"),
		UserID:    userID,
		Email:     email,
		Timestamp: time.Now(),
	}
}
