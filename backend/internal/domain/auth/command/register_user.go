package command

import (
	"10x-certification/internal/domain/auth/model"
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/auth/jwt"
	"10x-certification/internal/infrastructure/auth/password"
	"context"
)

// RegisterUserCommand represents the command to register a new user
type RegisterUserCommand struct {
	Email    string
	Password string
}

// NewRegisterUserCommand creates a new RegisterUserCommand
func NewRegisterUserCommand(email, password string) *RegisterUserCommand {
	return &RegisterUserCommand{
		Email:    email,
		Password: password,
	}
}

// RegisterUserHandler handles user registration
type RegisterUserHandler struct {
	userRepo       repository.UserRepository
	passwordHasher *password.Hasher
	jwtService     *jwt.TokenService
}

// NewRegisterUserHandler creates a new RegisterUserHandler
func NewRegisterUserHandler(
	userRepo repository.UserRepository,
	passwordHasher *password.Hasher,
	jwtService *jwt.TokenService,
) *RegisterUserHandler {
	return &RegisterUserHandler{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		jwtService:     jwtService,
	}
}

// Handle executes the register user command
func (h *RegisterUserHandler) Handle(ctx context.Context, cmd *RegisterUserCommand) (*model.User, error) {
	// TODO: Implement user registration logic
	// 1. Validate email format
	// 2. Check if user already exists
	// 3. Hash password
	// 4. Create user entity
	// 5. Save to repository
	// 6. Generate JWT token
	// 7. Publish UserRegistered event
	panic("not implemented")
}
