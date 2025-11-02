package routes

import (
	"golang-tutorial-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up authentication routes
type AuthRoutes struct {
	authHandler handlers.AuthHandler
}

// NewAuthRoutes creates a new auth routes instance
func NewAuthRoutes(authHandler handlers.AuthHandler) *AuthRoutes {
	return &AuthRoutes{
		authHandler: authHandler,
	}
}

// SetupRoutes configures all auth routes
func (ar *AuthRoutes) SetupRoutes(router *gin.Engine) {
	// Group all auth routes under /api/v1/auth
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/signup", ar.authHandler.Signup)
		authGroup.POST("/login", ar.authHandler.Login)
	}
}
