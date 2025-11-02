package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"school-management-backend/internal/handlers"
	"school-management-backend/internal/middlewares"
)

// SetupRouter initializes all the routes for the application
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(db)

	// Get JWT secret
	jwtSecret := getJWTSecret()

	// Setup routes
	setupHealthCheck(router, db)
	setupAuthRoutes(router, authHandler)
	setupProtectedRoutes(router, jwtSecret)

	return router
}

// setupHealthCheck configures the health check endpoint
func setupHealthCheck(router *gin.Engine, db *gorm.DB) {
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "error",
				"message": "Database connection issue",
				"error":   err.Error(),
			})
			return
		}

		// Try pinging the database to confirm connection
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status":  "error",
				"message": "Database ping failed",
				"error":   err.Error(),
			})
			return
		}

		// If everything works fine
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Server and database are healthy",
		})
	})
}

// setupAuthRoutes configures all authentication related routes
func setupAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/signup", authHandler.Signup)
		authGroup.POST("/login", authHandler.Login)
	}
}

// setupProtectedRoutes configures routes that require authentication
func setupProtectedRoutes(router *gin.Engine, jwtSecret string) {
	// All routes under this group will require a valid JWT token
	protected := router.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware(jwtSecret))
	
	// Example protected routes:
	// protected.GET("/profile", profileHandler.GetProfile)
	// protected.PUT("/profile", profileHandler.UpdateProfile)
}

// getJWTSecret retrieves the JWT secret from environment variables
func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// In production, you should fail fast if JWT_SECRET is not set
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	return secret
}
