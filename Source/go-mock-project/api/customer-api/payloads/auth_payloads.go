package customer_api_payloads

import (
	"time"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Email        string    `json:"email"`
	LastLoginAt  time.Time `json:"lastLoginAt"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpireIn     int64     `json:"expireIn"`
}
