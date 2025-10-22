package jwt

import (
	"errors"
)

// TokenValidator represents JWT token validator
type TokenValidator struct {
	tokenService *TokenService
}

// NewTokenValidator creates a new JWT token validator
func NewTokenValidator(tokenService *TokenService) *TokenValidator {
	return &TokenValidator{
		tokenService: tokenService,
	}
}

// ValidateToken validates a JWT token and returns claims
func (v *TokenValidator) ValidateToken(tokenString string) (*TokenClaims, error) {
	if tokenString == "" {
		return nil, errors.New("token is required")
	}

	claims, err := v.tokenService.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// ExtractTokenFromHeader extracts token from Authorization header
func (v *TokenValidator) ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	// TODO: Extract token from "Bearer <token>" format
	panic("not implemented")
}
