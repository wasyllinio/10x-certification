package query

import (
	"10x-certification/internal/domain/auth/model"
	"10x-certification/internal/domain/auth/repository"
	"context"
)

// GetUserByIDQuery represents the query to get user by ID
type GetUserByIDQuery struct {
	UserID string
}

// NewGetUserByIDQuery creates a new GetUserByIDQuery
func NewGetUserByIDQuery(userID string) *GetUserByIDQuery {
	return &GetUserByIDQuery{
		UserID: userID,
	}
}

// GetUserByIDHandler handles getting user by ID
type GetUserByIDHandler struct {
	userRepo repository.UserRepository
}

// NewGetUserByIDHandler creates a new GetUserByIDHandler
func NewGetUserByIDHandler(userRepo repository.UserRepository) *GetUserByIDHandler {
	return &GetUserByIDHandler{
		userRepo: userRepo,
	}
}

// Handle executes the get user by ID query
func (h *GetUserByIDHandler) Handle(ctx context.Context, query *GetUserByIDQuery) (*model.User, error) {
	// TODO: Implement get user by ID logic
	panic("not implemented")
}
