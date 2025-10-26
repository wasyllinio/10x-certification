package middleware

import (
	"10x-certification/internal/domain/auth/repository"
	"10x-certification/internal/infrastructure/auth/jwt"
	"10x-certification/internal/infrastructure/http/context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthMiddleware handles JWT authentication
func AuthMiddleware(jwtService *jwt.TokenService, userRepo repository.UserRepository) gin.HandlerFunc {
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

		// Parse authorization ID from string to UUID
		authorizationID, err := uuid.Parse(claims.AuthorizationID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization ID in token",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Fetch user from database
		user, err := userRepo.FindByID(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user not found",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Verify authorization ID matches
		if user.AuthorizationID != authorizationID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization credentials",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Verify role matches
		if string(user.Role) != claims.Role {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user role has changed, please login again",
				"code":  "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		// Set user context
		authCtx := &context.AuthContext{
			UserID:          userID,
			Email:           user.Email, // Use email from database
			Role:            string(user.Role),
			AuthorizationID: user.AuthorizationID,
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
