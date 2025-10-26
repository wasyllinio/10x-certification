package command

import (
	"10x-certification/internal/domain/auth/dto/request"
	"10x-certification/internal/domain/auth/model"
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/auth/jwt"
	"10x-certification/internal/infrastructure/auth/password"
	"10x-certification/internal/shared/errors"
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
	// 1. Find user by email
	userDB, err := h.userRepo.FindByEmail(ctx, cmd.Request.Email)
	if err != nil {
		return nil, "", errors.ErrInvalidCredentials // Nie ujawniamy, czy user istnieje
	}

	// 2. Verify password
	valid, err := h.passwordHasher.VerifyPassword(
		cmd.Request.Password,
		userDB.PasswordHash,
		userDB.PasswordSalt,
	)
	if err != nil || !valid {
		return nil, "", errors.ErrInvalidCredentials
	}

	// 3. Generate JWT token
	token, err := h.jwtService.GenerateToken(
		userDB.ID.String(),
		string(userDB.Role),
		userDB.AuthorizationID.String(),
	)
	if err != nil {
		return nil, "", err
	}

	// 4. Convert to domain model
	user := &model.User{
		ID:              userDB.ID,
		AuthorizationID: userDB.AuthorizationID,
		Email:           userDB.Email,
		PasswordHash:    userDB.PasswordHash,
		PasswordSalt:    userDB.PasswordSalt,
		Role:            model.UserRole(userDB.Role),
		CreatedAt:       userDB.CreatedAt,
		UpdatedAt:       userDB.UpdatedAt,
	}

	// 5. Publish UserLoggedIn event (opcjonalnie - jeśli event dispatcher jest gotowy)
	// event := events.NewUserLoggedIn(user.ID, user.Email)
	// eventDispatcher.Dispatch(event)

	return user, token, nil
}
