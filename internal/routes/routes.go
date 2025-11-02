package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"school-management-backend/internal/handler"
	"school-management-backend/internal/middlewares"
	"school-management-backend/internal/repository"
	"school-management-backend/internal/service"
)

// SetupRouter initializes all the routes for the application
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize repositories
	authService := service.NewAuthService(db)
	libraryRepo := repository.NewLibraryRepository(db)
	libraryService := service.NewLibraryService(libraryRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	libraryHandler := handler.NewLibraryHandler(libraryService)

	// Get JWT secret
	jwtSecret := getJWTSecret()

	// Setup public routes
	setupHealthCheck(router, db)
	setupAuthRoutes(router, authHandler)

	// Setup protected routes
	api := setupProtectedRoutes(router, jwtSecret)
	setupLibraryRoutes(api, libraryHandler)

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
func setupAuthRoutes(router *gin.Engine, authHandler *handler.AuthHandler) {
	authGroup := router.Group("/api/v1/auth")
	{
		authGroup.POST("/login", authHandler.Login)
	}
}

// setupProtectedRoutes configures routes that require authentication
// and returns the API router group
func setupProtectedRoutes(router *gin.Engine, jwtSecret string) *gin.RouterGroup {
	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware(jwtSecret))
	return api
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
