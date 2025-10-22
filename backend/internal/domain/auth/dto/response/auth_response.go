package response

// AuthResponse represents the authentication response containing JWT token
type AuthResponse struct {
	Token string `json:"token"`
}

// NewAuthResponse creates a new AuthResponse
func NewAuthResponse(token string) *AuthResponse {
	return &AuthResponse{
		Token: token,
	}
}
