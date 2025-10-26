package command

import (
	"10x-certification/internal/domain/auth/dto/request"
	"10x-certification/internal/domain/auth/model"
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/auth/jwt"
	"10x-certification/internal/infrastructure/auth/password"
	"context"
)

// LoginUserCommand represents the command to login a user
type LoginUserCommand struct {
	Request *request.LoginRequest
}

// NewLoginUserCommand creates a new LoginUserCommand
func NewLoginUserCommand(req *request.LoginRequest) *LoginUserCommand {
	return &LoginUserCommand{
		Request: req,
	}
}

// LoginUserHandler handles user login
type LoginUserHandler struct {
	userRepo       repository.UserRepository
	passwordHasher *password.Hasher
	jwtService     *jwt.TokenService
}

// NewLoginUserHandler creates a new LoginUserHandler
func NewLoginUserHandler(
	userRepo repository.UserRepository,
	passwordHasher *password.Hasher,
	jwtService *jwt.TokenService,
) *LoginUserHandler {
	return &LoginUserHandler{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		jwtService:     jwtService,
	}
}

// Handle executes the login user command
func (h *LoginUserHandler) Handle(ctx context.Context, cmd *LoginUserCommand) (*model.User, string, error) {
	// TODO: Implement user login logic
	// 1. Find user by email
	// 2. Verify password
	// 3. Generate JWT token
	// 4. Publish UserLoggedIn event
	panic("not implemented")
}
