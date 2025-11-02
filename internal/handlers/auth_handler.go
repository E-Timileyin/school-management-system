package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthHandler defines the interface for authentication handlers
type AuthHandler interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
}

type authHandler struct {
	// Add any dependencies here (e.g., user service, logger)
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

// hashPassword hashes the given password using bcrypt
func (h *authHandler) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Signup handles user registration
func (h *authHandler) Signup(c *gin.Context) {
	// Define request body structure
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	// Parse and validate JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	if _, err := h.hashPassword(input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// TODO: Save user to database
	// user := model.User{
	// 	Email:    input.Email,
	// 	Password: hashedPassword,
	// }
	// if err := user.Create(); err != nil { ... }

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"email":   input.Email,
	})
}

// Login function

func (h *authHandler) Login(c *gin.Context) {
	// parse input json
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// parse and validate json
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Verify credentials
	// user := model.User{
	// 	Email:    input.Email,
	// 	Password: input.Password,
	// }
	// if err := user.Validate(); err != nil { ... }

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"email":   input.Email,
	})
}
