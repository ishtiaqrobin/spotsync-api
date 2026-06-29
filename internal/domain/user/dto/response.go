package dto

// Response represents a safe user response (no password exposed)
type Response struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// LoginResponse represents the response after successful login
type LoginResponse struct {
	Token string        `json:"token"`
	User  LoginUserInfo `json:"user"`
}

// LoginUserInfo represents the user info returned in login response
type LoginUserInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
