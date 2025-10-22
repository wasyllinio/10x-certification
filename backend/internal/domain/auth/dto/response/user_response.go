package response

import (
	"time"

	"github.com/google/uuid"
)

// UserResponse represents user data response
type UserResponse struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	Role            string    `json:"role"`
	AuthorizationID uuid.UUID `json:"authorization_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// NewUserResponse creates a new UserResponse from domain User
func NewUserResponse(id, authorizationID uuid.UUID, email, role string, createdAt, updatedAt time.Time) *UserResponse {
	return &UserResponse{
		ID:              id,
		Email:           email,
		Role:            role,
		AuthorizationID: authorizationID,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
