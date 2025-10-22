package query

import (
	"10x-certification/internal/domain/auth/model"
	"10x-certification/internal/domain/auth/repository"
	"context"
)

// GetUserByEmailQuery represents the query to get user by email
type GetUserByEmailQuery struct {
	Email string
}

// NewGetUserByEmailQuery creates a new GetUserByEmailQuery
func NewGetUserByEmailQuery(email string) *GetUserByEmailQuery {
	return &GetUserByEmailQuery{
		Email: email,
	}
}

// GetUserByEmailHandler handles getting user by email
type GetUserByEmailHandler struct {
	userRepo repository.UserRepository
}

// NewGetUserByEmailHandler creates a new GetUserByEmailHandler
func NewGetUserByEmailHandler(userRepo repository.UserRepository) *GetUserByEmailHandler {
	return &GetUserByEmailHandler{
		userRepo: userRepo,
	}
}

// Handle executes the get user by email query
func (h *GetUserByEmailHandler) Handle(ctx context.Context, query *GetUserByEmailQuery) (*model.User, error) {
	// TODO: Implement get user by email logic
	panic("not implemented")
}
