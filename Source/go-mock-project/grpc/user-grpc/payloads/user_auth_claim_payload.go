package user_grpc_payloads

import "github.com/golang-jwt/jwt"

type UserAuthClaims struct {
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	jwt.StandardClaims
}
