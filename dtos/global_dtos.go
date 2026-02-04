package dtos

import "time"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
type IDResponse struct {
	ID uint `json:"id"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	Role      string `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}
