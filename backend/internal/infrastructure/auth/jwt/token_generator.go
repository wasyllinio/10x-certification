package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenService represents JWT token service
type TokenService struct {
	secretKey string
}

// NewTokenService creates a new JWT token service
func NewTokenService(secretKey string) *TokenService {
	return &TokenService{
		secretKey: secretKey,
	}
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID          string `json:"user_id"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	AuthorizationID string `json:"authorization_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for user
func (s *TokenService) GenerateToken(userID, email, role, authorizationID string) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		UserID:          userID,
		Email:           email,
		Role:            role,
		AuthorizationID: authorizationID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)), // 24 hours
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

// ValidateToken validates a JWT token
func (s *TokenService) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken generates a new token with extended expiration
func (s *TokenService) RefreshToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Generate new token with extended expiration
	return s.GenerateToken(claims.UserID, claims.Email, claims.Role, claims.AuthorizationID)
}
