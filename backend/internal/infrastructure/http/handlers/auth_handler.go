package handlers

import (
	"10x-certification/internal/domain/auth/command"
	"10x-certification/internal/domain/auth/dto/request"
	"10x-certification/internal/domain/auth/dto/response"
	"10x-certification/internal/domain/auth/query"
	"10x-certification/internal/infrastructure/http/dto"
	"10x-certification/internal/shared/errors"
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
	// 1. Bind and validate request
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := dto.FormatValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"code":    "VALIDATION_ERROR",
			"details": validationErrors,
		})
		return
	}

	// 2. Create command
	cmd := command.NewRegisterUserCommand(&req)

	// 3. Execute command
	_, err := h.registerUserHandler.Handle(c.Request.Context(), cmd)
	if err != nil {
		httpErr := errors.MapDomainErrorToHTTP(err)
		c.JSON(httpErr.StatusCode, gin.H{
			"error": httpErr.Message,
			"code":  httpErr.Code,
		})
		return
	}

	// 4. Return success response
	c.JSON(http.StatusCreated, gin.H{})
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	// 1. Bind and validate request
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := dto.FormatValidationErrors(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "validation failed",
			"code":    "VALIDATION_ERROR",
			"details": validationErrors,
		})
		return
	}

	// 2. Create command
	cmd := command.NewLoginUserCommand(&req)

	// 3. Execute command
	_, token, err := h.loginUserHandler.Handle(c.Request.Context(), cmd)
	if err != nil {
		httpErr := errors.MapDomainErrorToHTTP(err)
		c.JSON(httpErr.StatusCode, gin.H{
			"error": httpErr.Message,
			"code":  httpErr.Code,
		})
		return
	}

	// 4. Return JWT token
	resp := response.NewAuthResponse(token)
	c.JSON(http.StatusOK, resp)
}

// GetUser handles getting user by ID
func (h *AuthHandler) GetUser(c *gin.Context) {
	// TODO: Implement get user HTTP handler
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
