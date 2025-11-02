package utils

import "github.com/golang-jwt/jwt/v5"

// JWTClaims represents the claims to be included in the JWT token
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
