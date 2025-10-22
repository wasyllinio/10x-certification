package handlers

import (
	"10x-certification/internal/domain/auth/command"
	"10x-certification/internal/domain/auth/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication HTTP requests
type AuthHandler struct {
	registerUserHandler   *command.RegisterUserHandler
	loginUserHandler      *command.LoginUserHandler
	getUserByIDHandler    *query.GetUserByIDHandler
	getUserByEmailHandler *query.GetUserByEmailHandler
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(
	registerUserHandler *command.RegisterUserHandler,
	loginUserHandler *command.LoginUserHandler,
	getUserByIDHandler *query.GetUserByIDHandler,
	getUserByEmailHandler *query.GetUserByEmailHandler,
) *AuthHandler {
	return &AuthHandler{
		registerUserHandler:   registerUserHandler,
		loginUserHandler:      loginUserHandler,
		getUserByIDHandler:    getUserByIDHandler,
		getUserByEmailHandler: getUserByEmailHandler,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	// TODO: Implement user registration HTTP handler
	// 1. Bind JSON request to DTO
	// 2. Create command
	// 3. Call command handler
	// 4. Return response
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	// TODO: Implement user login HTTP handler
	// 1. Bind JSON request to DTO
	// 2. Create command
	// 3. Call command handler
	// 4. Return JWT token
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}

// GetUser handles getting user by ID
func (h *AuthHandler) GetUser(c *gin.Context) {
	// TODO: Implement get user HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
