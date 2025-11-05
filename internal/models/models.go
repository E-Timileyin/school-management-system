package models

import (
	"errors"

	"github.com/E-Timileyin/school-management-system/internal/utils"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Role      string `gorm:"not null"` // admin, teacher, student, parent
}

// SetPassword hashes the password and sets it on the user
func (u *User) SetPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

// CheckPassword verifies if the provided password matches the hashed password
func (u *User) CheckPassword(password string) error {
	return utils.CheckPassword(u.Password, password)
}

// Student represents a student in the system
type Student struct {
	gorm.Model
	UserID      uint   `gorm:"uniqueIndex;not null"`
	User        User   `gorm:"foreignKey:UserID"`
	DateOfBirth string `gorm:"type:date"`
	Address     string
	Phone       string
}

// Teacher represents a teacher in the system
type Teacher struct {
	gorm.Model
	UserID  uint   `gorm:"uniqueIndex;not null"`
	User    User   `gorm:"foreignKey:UserID"`
	Subject string `gorm:"not null"`
	Phone   string
}

// Course represents a course in the system
type Course struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Code        string `gorm:"uniqueIndex;not null"`
	Description string
	TeacherID   uint    `gorm:"not null"` // Reference to Teacher
	Teacher     Teacher `gorm:"foreignKey:TeacherID"`
}

// Enrollment represents a student's enrollment in a course
type Enrollment struct {
	gorm.Model
	StudentID uint    `gorm:"not null"`
	CourseID  uint    `gorm:"not null"`
	Grade     *string // Nullable grade
}

// Set up table names for all models
func (User) TableName() string {
	return "users"
}

func (Student) TableName() string {
	return "students"
}

func (Teacher) TableName() string {
	return "teachers"
}

func (Course) TableName() string {
	return "courses"
}

func (Enrollment) TableName() string {
	return "enrollments"
}
