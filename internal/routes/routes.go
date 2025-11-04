package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/E-Timileyin/school-management-system/internal/handler"
	"github.com/E-Timileyin/school-management-system/internal/middlewares"
	"github.com/E-Timileyin/school-management-system/internal/repository"
	"github.com/E-Timileyin/school-management-system/internal/service"
)

// SetupRouter initializes all the routes for the application
func SetupRouter(db *gorm.DB) *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)
	// authRepo is not needed as userRepo handles authentication
	libraryRepo := repository.NewLibraryRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	courseService := service.NewCourseService(courseRepo)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo)
	// authService is not needed as userService handles authentication
	libraryService := service.NewLibraryService(libraryRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	courseHandler := handler.NewCourseHandler(courseService, enrollmentService)
	// authHandler is not needed as userHandler handles authentication
	libraryHandler := handler.NewLibraryHandler(libraryService)
	adminHandler := handler.NewAdminHandler(userService, courseService)

	// Get JWT secret
	jwtSecret := getJWTSecret()

	// ====== Public Routes ======
	setupHealthCheck(router, db)
	// Auth routes are handled by userHandler
	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)

	// ====== Protected API Routes ======
	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware(jwtSecret))
	{
		// User profile routes
		setupUserRoutes(api, userHandler)

		// Library routes
		setupLibraryRoutes(api, libraryHandler)

		// Course routes
		setupCourseRoutes(api, courseHandler)
	}

	// ====== Admin Routes ======
	admin := router.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(jwtSecret))
	admin.Use(middlewares.AdminMiddleware()) // Only users with admin role can access
	{
		setupAdminRoutes(admin, adminHandler)
	}

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

// setupUserRoutes configures user profile related routes
func setupUserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	users := router.Group("/users")
	{
		users.GET("/me", userHandler.GetProfile)
		users.PUT("/me", userHandler.UpdateProfile)
		users.PUT("/password", userHandler.ChangePassword)
	}
}

// setupAdminRoutes configures admin management routes
func setupAdminRoutes(router *gin.RouterGroup, adminHandler *handler.AdminHandler) {
	// User management
	users := router.Group("/users")
	{
		users.GET("", adminHandler.GetAllUsers)
		users.POST("", adminHandler.CreateUser)
		users.GET("/:id", adminHandler.GetUserByID)
		users.PUT("/:id", adminHandler.UpdateUser)
		users.DELETE("/:id", adminHandler.DeleteUser)
	}

	// Course management
	courses := router.Group("/courses")
	{
		courses.POST("", adminHandler.CreateCourse)
		courses.PUT("/:id", adminHandler.UpdateCourse)
		courses.DELETE("/:id", adminHandler.DeleteCourse)
		courses.POST("/:id/enrollments", adminHandler.EnrollStudent)
		courses.DELETE("/:courseId/enrollments/:studentId", adminHandler.RemoveEnrollment)
	}
}

// setupCourseRoutes configures course related routes
func setupCourseRoutes(router *gin.RouterGroup, courseHandler *handler.CourseHandler) {
	courses := router.Group("/courses")
	{
		courses.GET("", courseHandler.GetAllCourses)
		courses.GET("/:id", courseHandler.GetCourseByID)
		courses.GET("/:id/students", courseHandler.GetCourseStudents)
	}

	// Student enrollments
	enrollments := router.Group("/enrollments")
	{
		enrollments.GET("", courseHandler.GetMyEnrollments)
		enrollments.POST("/:courseId", courseHandler.EnrollInCourse)
		enrollments.DELETE("/:enrollmentId", courseHandler.WithdrawFromCourse)
	}
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
