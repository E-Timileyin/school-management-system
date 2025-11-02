package model

import (
	"errors"
	"time"

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
	Email          string    `gorm:"uniqueIndex;not null;size:255"`
	PasswordHash   string    `gorm:"not null;size:255"`
	LastLogin      time.Time `gorm:"default:null"`
	Role           UserRole  `gorm:"type:user_role;default:'student'"`
	IsActive       bool      `gorm:"default:false"`
	EmailVerified  bool      `gorm:"default:false"`
	
	// Profile Information
	FirstName      string    `gorm:"size:100;not null"`
	LastName       string    `gorm:"size:100;not null"`
	DateOfBirth    time.Time `gorm:"type:date"`
	PhoneNumber    string    `gorm:"size:20"`
	ProfilePicture string    `gorm:"size:255"`
	
	// Account Status
	IsSuspended   bool      `gorm:"default:false"`
	SuspendedAt   time.Time `gorm:"default:null"`
	SuspendedBy   uint      // ID of the admin who suspended this user
	LastPasswordChange time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// BeforeCreate is a GORM hook that runs before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.LastPasswordChange = time.Now()
	return nil
}

// SetPassword hashes the password and stores it in the PasswordHash field
func (u *User) SetPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	u.PasswordHash = string(hashedPassword)
	u.LastPasswordChange = time.Now()
	return nil
}

// CheckPassword compares the provided password with the stored hash
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
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
