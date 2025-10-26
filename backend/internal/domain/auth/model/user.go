package model

import (
	"time"

	"github.com/google/uuid"
)

// UserRole represents user role enum
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleOwner UserRole = "owner"
)

// User represents the user aggregate root
type User struct {
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Email           string
	PasswordHash    string
	PasswordSalt    string
	Role            UserRole
	ID              uuid.UUID
	AuthorizationID uuid.UUID
}

// NewUser creates a new user entity
func NewUser(email string, passwordHash, passwordSalt string, role UserRole) *User {
	return &User{
		Email:           email,
		PasswordHash:    passwordHash,
		PasswordSalt:    passwordSalt,
		Role:            role,
	}
}

// IsAdmin checks if user has admin role
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// IsOwner checks if user has owner role
func (u *User) IsOwner() bool {
	return u.Role == RoleOwner
}

// CanAccess checks if user can access resource owned by ownerID
func (u *User) CanAccess(ownerID uuid.UUID) bool {
	return u.IsAdmin() || u.ID == ownerID
}
