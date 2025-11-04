// internal/handler/user_handler.go
package handler

import (
	"net/http"

	"github.com/E-Timileyin/school-management-system/internal/models"
	"github.com/E-Timileyin/school-management-system/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	userService *service.UserService
}

// internal/handler/user_handler.go
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Login handles user login
func (h *UserHandler) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.userService.GetUserByEmail(loginData.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if err := user.CheckPassword(loginData.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// TODO: Generate JWT token and return it
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"user":    user,
	})
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var registerData struct {
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Role      string `json:"role" binding:"required,oneof=admin teacher student parent"`
	}

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	user := &models.User{
		Email:     registerData.Email,
		FirstName: registerData.FirstName,
		LastName:  registerData.LastName,
		Role:      registerData.Role,
	}

	if err := user.SetPassword(registerData.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password: " + err.Error()})
		return
	}

	if err := h.userService.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// Don't return the password hash in the response
	user.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	currentUser := user.(*models.User)

	var updateData struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	currentUser.FirstName = updateData.FirstName
	currentUser.LastName = updateData.LastName
	currentUser.Email = updateData.Email

	if err := h.userService.UpdateUser(currentUser); err != nil {
		c.JSON(500, gin.H{"error": "failed to update profile"})
		return
	}

	c.JSON(200, currentUser)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	currentUser := user.(*models.User)

	var passwordData struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&passwordData); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if err := currentUser.CheckPassword(passwordData.CurrentPassword); err != nil {
		c.JSON(400, gin.H{"error": "current password is incorrect"})
		return
	}

	if err := currentUser.SetPassword(passwordData.NewPassword); err != nil {
		c.JSON(400, gin.H{"error": "invalid new password"})
		return
	}

	if err := h.userService.UpdateUser(currentUser); err != nil {
		c.JSON(500, gin.H{"error": "failed to update password"})
		return
	}

	c.JSON(200, gin.H{"message": "password updated successfully"})
}
