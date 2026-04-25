package dto

type AuthResponse struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
