package request

// LoginRequest represents the request to login a user
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// NewLoginRequest creates a new LoginRequest
func NewLoginRequest(email, password string) *LoginRequest {
	return &LoginRequest{
		Email:    email,
		Password: password,
	}
}
