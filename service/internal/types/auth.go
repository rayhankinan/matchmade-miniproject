package types

// RegisterRequest represents the data needed for a user to register
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginRequest represents the data needed for a user to log in
type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required,identifier"`
	Password   string `json:"password" validate:"required"`
}
