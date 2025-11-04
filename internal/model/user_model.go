package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRole represents the role of a user in the system
type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleTeacher  UserRole = "teacher"
	RoleStudent  UserRole = "student"
	RoleParent   UserRole = "parent"
	RoleGuardian UserRole = "guardian"
)

type User struct {
	gorm.Model

	// Authentication
	Email     string   `gorm:"uniqueIndex;not null"`
	Password  string   `gorm:"not null"` // Maps to 'password' column in database
	FirstName string   `gorm:"not null"`
	LastName  string   `gorm:"not null"`
	Role      UserRole `gorm:"not null"`
}

// BeforeCreate is a GORM hook that runs before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Role == "" {
		u.Role = RoleStudent
	}
	return nil
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies the password
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// IsAdmin checks if the user has admin role
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// CanAccess checks if user has permission to access a resource
func (u *User) CanAccess(requiredRole UserRole) bool {
	if u.IsAdmin() {
		return true
	}
	return u.Role == requiredRole
}
