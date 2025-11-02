package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"school-management-backend/internal/model"
	"school-management-backend/internal/utils"
)

// AuthHandler handles signup and login requests
type AuthHandler struct {
	DB *gorm.DB
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// Create a new AuthHandler (so we can use the database)
func (h *AuthHandler) Signup(c *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existingUser model.User
	if err := h.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Create new user with hashed password
	newUser := model.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		IsActive:  true,              // Or set to false if email verification is required
		Role:      model.RoleStudent, // Default role
	}

	// Set password using the helper method
	if err := newUser.SetPassword(input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save user to database
	if err := h.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user: " + err.Error()})
		return
	}

	// TODO: Send verification email if needed

	// Return success response (don't return sensitive data)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully. Please check your email to verify your account.",
		"user": gin.H{
			"id":        newUser.ID,
			"email":     newUser.Email,
			"firstName": newUser.FirstName,
			"lastName":  newUser.LastName,
		},
	})
}

// ====================// Login handles user authentication
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user model.User
	if err := h.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		// Don't reveal if user exists or not for security
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if account is active
	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"error": "Account is not active. Please verify your email or contact support."})
		return
	}

	// Check password using the helper method
	if err := user.CheckPassword(input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token (you should get the secret from config)
	jwtSecret := "your-secret-key" // TODO: Get from config
	token, err := utils.GenerateToken(user, jwtSecret, 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Update last login time
	user.LastLogin = time.Now()
	h.DB.Save(&user)

	// Return token and user info (don't include sensitive data)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"role":      user.Role,
		},
	})
}
