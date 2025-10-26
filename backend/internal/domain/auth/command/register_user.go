package command

import (
	"10x-certification/internal/domain/auth/dto/request"
	"10x-certification/internal/domain/auth/model"
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/auth/password"
	"10x-certification/internal/infrastructure/persistence/models"
	"10x-certification/internal/shared/errors"
	"context"
)

// RegisterUserCommand represents the command to register a new user
type RegisterUserCommand struct {
	Request *request.RegisterRequest
}

// NewRegisterUserCommand creates a new RegisterUserCommand
func NewRegisterUserCommand(req *request.RegisterRequest) *RegisterUserCommand {
	return &RegisterUserCommand{
		Request: req,
	}
}

// RegisterUserHandler handles user registration
type RegisterUserHandler struct {
	userRepo       repository.UserRepository
	passwordHasher *password.Hasher
}

// NewRegisterUserHandler creates a new RegisterUserHandler
func NewRegisterUserHandler(
	userRepo repository.UserRepository,
	passwordHasher *password.Hasher,
) *RegisterUserHandler {
	return &RegisterUserHandler{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

// Handle executes the register user command
func (h *RegisterUserHandler) Handle(ctx context.Context, cmd *RegisterUserCommand) (*model.User, error) {
	// 1. Check if user already exists
	exists, err := h.userRepo.ExistsByEmail(ctx, cmd.Request.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.ErrUserAlreadyExists
	}

	// 2. Hash password
	passwordHash, passwordSalt, err := h.passwordHasher.HashPassword(cmd.Request.Password)
	if err != nil {
		return nil, err
	}

	// 3. Create domain User entity
	user := model.NewUser(cmd.Request.Email, passwordHash, passwordSalt, model.RoleOwner)

	// 4. Convert to database model
	userDB := &models.UserDB{
		ID:              user.ID,
		Email:           user.Email,
		PasswordHash:    user.PasswordHash,
		PasswordSalt:    user.PasswordSalt,
		Role:            models.UserRole(user.Role),
		AuthorizationID: user.AuthorizationID.String(),
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}

	// 5. Save to repository
	if err := h.userRepo.Save(ctx, userDB); err != nil {
		return nil, err
	}

	// 6. Return domain user
	return user, nil
}
