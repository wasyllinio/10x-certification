package errors

import (
	"net/http"
)

// HTTPError represents an HTTP-specific error
type HTTPError struct {
	Message    string
	Code       string
	StatusCode int
}

// Error implements the error interface
func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTP error
func NewHTTPError(statusCode int, message, code string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
		Code:       code,
	}
}

// Common HTTP error responses
var (
	ErrBadRequest          = NewHTTPError(http.StatusBadRequest, "bad request", "BAD_REQUEST")
	ErrUnauthorized        = NewHTTPError(http.StatusUnauthorized, "unauthorized", "UNAUTHORIZED")
	ErrForbidden           = NewHTTPError(http.StatusForbidden, "forbidden", "FORBIDDEN")
	ErrNotFound            = NewHTTPError(http.StatusNotFound, "not found", "NOT_FOUND")
	ErrMethodNotAllowed    = NewHTTPError(http.StatusMethodNotAllowed, "method not allowed", "METHOD_NOT_ALLOWED")
	ErrConflict            = NewHTTPError(http.StatusConflict, "conflict", "CONFLICT")
	ErrUnprocessableEntity = NewHTTPError(http.StatusUnprocessableEntity, "unprocessable entity", "UNPROCESSABLE_ENTITY")
	ErrInternalServerError = NewHTTPError(http.StatusInternalServerError, "internal server error", "INTERNAL_SERVER_ERROR")
)
