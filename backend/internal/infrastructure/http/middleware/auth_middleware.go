package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


// ErrorMiddleware handles HTTP errors
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Handle errors that occurred during request processing
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			// TODO: Map domain errors to HTTP errors
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	}
}

// AuthMiddleware handles JWT authentication
func AuthMiddleware(jwtService interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT authentication middleware
		// 1. Extract token from Authorization header
		// 2. Validate token
		// 3. Set user context
		c.Next()
	}
}

// AuthorizationMiddleware handles role-based authorization
func AuthorizationMiddleware(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement role-based authorization middleware
		// 1. Get user from context
		// 2. Check if user role is allowed
		// 3. Return 403 if not authorized
		c.Next()
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
