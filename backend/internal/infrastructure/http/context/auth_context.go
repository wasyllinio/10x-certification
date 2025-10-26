package context

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// contextKey is a custom type to avoid collisions with other packages
type contextKey string

const (
	authContextKey contextKey = "auth_context"
)

// AuthContext holds all authentication-related data from JWT claims
type AuthContext struct {
	Email           string
	Role            string
	AuthorizationID string
	UserID          uuid.UUID
}

// SetAuthContext stores authentication context in Gin context
func SetAuthContext(c *gin.Context, auth *AuthContext) {
	c.Set(string(authContextKey), auth)
}

// GetAuthContext retrieves authentication context from Gin context
func GetAuthContext(c *gin.Context) (*AuthContext, error) {
	value, exists := c.Get(string(authContextKey))
	if !exists {
		return nil, errors.New("authentication context not found")
	}

	auth, ok := value.(*AuthContext)
	if !ok {
		return nil, errors.New("invalid authentication context type")
	}

	return auth, nil
}

// GetUserID retrieves user ID from authentication context
func GetUserID(c *gin.Context) (uuid.UUID, error) {
	auth, err := GetAuthContext(c)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get user ID: %w", err)
	}
	return auth.UserID, nil
}

// GetUserEmail retrieves user email from authentication context
func GetUserEmail(c *gin.Context) (string, error) {
	auth, err := GetAuthContext(c)
	if err != nil {
		return "", fmt.Errorf("failed to get user email: %w", err)
	}
	return auth.Email, nil
}

// GetUserRole retrieves user role from authentication context
func GetUserRole(c *gin.Context) (string, error) {
	auth, err := GetAuthContext(c)
	if err != nil {
		return "", fmt.Errorf("failed to get user role: %w", err)
	}
	return auth.Role, nil
}

// GetAuthorizationID retrieves authorization ID from authentication context
func GetAuthorizationID(c *gin.Context) (string, error) {
	auth, err := GetAuthContext(c)
	if err != nil {
		return "", fmt.Errorf("failed to get authorization ID: %w", err)
	}
	return auth.AuthorizationID, nil
}
