package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string, jwtSecret string) (*JWTClaims, error) {
	// token validation here
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*JWTClaims), nil
}
