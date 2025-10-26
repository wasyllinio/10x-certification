package middleware

import (
	"10x-certification/internal/shared/errors"

	"github.com/gin-gonic/gin"
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
