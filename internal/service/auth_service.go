// Package service handles business logic and authentication
package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/E-Timileyin/school-management-system/internal/model"
)

// AuthService handles user authentication and authorization
type AuthService struct {
	db *gorm.DB // Database connection
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// Login verifies user credentials and returns a JWT token if successful
// Returns user details and token if login is successful
func (s *AuthService) Login(email, password string) (*model.User, string, error) {
	// Find user by email
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	// Check if password matches
	if err := user.CheckPassword(password); err != nil {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	// Save the user (without LastLogin since it was removed from the model)
	s.db.Save(&user)

	// Generate JWT token for authenticated user
	token, err := s.generateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	return &user, token, nil
}

// generateJWT creates a signed JWT token for authenticated users
// Token includes user ID, role, and expiration time
func (s *AuthService) generateJWT(userID uint, role string) (string, error) {
	// TODO: Move secret key to environment variable
	jwtSecret := []byte("your-secret-key")

	// Set token claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                     // User identifier
		"role":    role,                       // User role (admin, teacher, student)
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign token with secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

// Signup creates a new user account
// Returns the created user or an error if registration fails
func (s *AuthService) Signup(email, password, firstName, lastName, role string) (*model.User, error) {
	// Check if email is already registered
	var existingUser model.User
	if err := s.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("email already registered")
	}

	// Create new user with minimal required fields
	user := &model.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Role:      model.UserRole(role),
	}

	// Set the password (this will hash it)
	if err := user.SetPassword(password); err != nil {
		return nil, fmt.Errorf("failed to set password: %v", err)
	}

	// Save user to database
	if err := s.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

// Register handles user registration with a pre-created user object
// Hashes password before saving to database
func (s *AuthService) Register(user *model.User) error {
	// Check if email is already registered
	var existingUser model.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return fmt.Errorf("email already registered")
	}

	// Set default role if not provided
	if user.Role == "" {
		user.Role = model.RoleStudent // Default role
	}

	// Save user to database (password should be set using SetPassword before calling Register)
	if user.Password == "" {
		return fmt.Errorf("password not set")
	}
	
	// Ensure the role is set
	if user.Role == "" {
		user.Role = model.RoleStudent
	}

	return s.db.Create(user).Error
}
