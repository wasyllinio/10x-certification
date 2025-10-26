package request

// RegisterRequest represents the request to register a new user
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// NewRegisterRequest creates a new RegisterRequest
func NewRegisterRequest(email, password string) *RegisterRequest {
	return &RegisterRequest{
		Email:    email,
		Password: password,
	}
}
