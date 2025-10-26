package middleware

import (
	"10x-certification/internal/infrastructure/auth/jwt"
	"10x-certification/internal/infrastructure/http/context"
	"10x-certification/internal/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ErrorMiddleware handles HTTP errors
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Handle errors that occurred during request processing
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// Map domain errors to HTTP errors
			httpErr := errors.MapDomainErrorToHTTP(err)
			c.JSON(httpErr.StatusCode, gin.H{
				"error": httpErr.Message,
				"code":  httpErr.Code,
			})
		}
	}
}

// AuthMiddleware handles JWT authentication
func AuthMiddleware(jwtService *jwt.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>" format
		tokenValidator := jwt.NewTokenValidator(jwtService)
		token, err := tokenValidator.ExtractTokenFromHeader(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header format",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Validate token
		claims, err := tokenValidator.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Parse user ID from string to UUID
		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user ID in token",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Set user context
		authCtx := &context.AuthContext{
			UserID:          userID,
			Email:           claims.Email,
			Role:            claims.Role,
			AuthorizationID: claims.AuthorizationID,
		}
		context.SetAuthContext(c, authCtx)

		c.Next()
	}
}

// AuthorizationMiddleware handles role-based authorization
func AuthorizationMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context
		userRole, err := context.GetUserRole(c)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "user role not found in context",
				"code":  "FORBIDDEN",
			})
			c.Abort()
			return
		}

		// Check if role is in allowed roles
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "access denied",
			"code":  "FORBIDDEN",
		})
		c.Abort()
	}
}

// CORSMiddleware handles CORS headers
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
