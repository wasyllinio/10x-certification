package errors

import (
	"errors"
	"net/http"
)

// DomainError represents a domain-specific error
type DomainError struct {
	Code    string
	Message string
	Cause   error
}

// Error implements the error interface
func (e *DomainError) Error() string {
	if e.Cause != nil {
		return e.Message + ": " + e.Cause.Error()
	}
	return e.Message
}

// Unwrap returns the underlying error
func (e *DomainError) Unwrap() error {
	return e.Cause
}

// NewDomainError creates a new domain error
func NewDomainError(code, message string, cause error) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// Common domain error codes
const (
	ErrCodeValidation    = "VALIDATION_ERROR"
	ErrCodeNotFound      = "NOT_FOUND"
	ErrCodeAlreadyExists = "ALREADY_EXISTS"
	ErrCodeUnauthorized  = "UNAUTHORIZED"
	ErrCodeForbidden     = "FORBIDDEN"
	ErrCodeConflict      = "CONFLICT"
	ErrCodeInternal      = "INTERNAL_ERROR"
)

// Common domain errors
var (
	ErrUserNotFound       = NewDomainError(ErrCodeNotFound, "user not found", nil)
	ErrUserAlreadyExists  = NewDomainError(ErrCodeAlreadyExists, "user already exists", nil)
	ErrChargerNotFound    = NewDomainError(ErrCodeNotFound, "charger not found", nil)
	ErrLocationNotFound   = NewDomainError(ErrCodeNotFound, "location not found", nil)
	ErrInvalidCredentials = NewDomainError(ErrCodeUnauthorized, "invalid credentials", nil)
	ErrAccessDenied       = NewDomainError(ErrCodeForbidden, "access denied", nil)
	ErrValidationFailed   = NewDomainError(ErrCodeValidation, "validation failed", nil)
)

// MapDomainErrorToHTTP maps domain errors to HTTP errors
func MapDomainErrorToHTTP(err error) *HTTPError {
	var domainErr *DomainError
	if errors.As(err, &domainErr) {
		switch domainErr.Code {
		case ErrCodeNotFound:
			return NewHTTPError(http.StatusNotFound, domainErr.Message, domainErr.Code)
		case ErrCodeAlreadyExists:
			return NewHTTPError(http.StatusConflict, domainErr.Message, domainErr.Code)
		case ErrCodeUnauthorized:
			return NewHTTPError(http.StatusUnauthorized, domainErr.Message, domainErr.Code)
		case ErrCodeForbidden:
			return NewHTTPError(http.StatusForbidden, domainErr.Message, domainErr.Code)
		case ErrCodeValidation:
			return NewHTTPError(http.StatusBadRequest, domainErr.Message, domainErr.Code)
		case ErrCodeConflict:
			return NewHTTPError(http.StatusConflict, domainErr.Message, domainErr.Code)
		default:
			return NewHTTPError(http.StatusInternalServerError, "internal server error", ErrCodeInternal)
		}
	}

	return NewHTTPError(http.StatusInternalServerError, "internal server error", ErrCodeInternal)
}
