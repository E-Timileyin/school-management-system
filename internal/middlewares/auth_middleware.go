package middlewares

import (
	"net/http"
	"github.com/E-Timileyin/school-management-system/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	// This file will hold the function that checks if a user has a valid token before allowing access to protected routes.
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Check if header exists and starts with Bearer
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			c.Abort() // stop the request
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token with the provided JWT secret
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Store user info in Gin context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
