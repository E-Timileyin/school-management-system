package utils

import (
	"time"

	"github.com/E-Timileyin/school-management-system/internal/model"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a new JWT token for the given user
func GenerateToken(user model.User, jwtSecret string, expirationTime time.Duration) (string, error) {
	expiration := time.Now().Add(expirationTime)

	// Create the JWT claims
	claims := &JWTClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "school-management-system",
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
