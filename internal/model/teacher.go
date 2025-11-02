package model

import "time"

type TeacherStatus string

const (
	TeacherStatusActive   TeacherStatus = "active"
	TeacherStatusInactive TeacherStatus = "inactive"
	TeacherStatusOnLeave  TeacherStatus = "on_leave"
)

type Teacher struct {
	Base
	UserID         uint          `gorm:"not null;uniqueIndex" json:"user_id"`
	EmployeeID     string        `gorm:"size:50;unique;not null" json:"employee_id"`
	JoiningDate    time.Time     `gorm:"not null" json:"joining_date"`
	Qualification  string        `gorm:"size:255" json:"qualification,omitempty"`
	Experience     string        `gorm:"size:100" json:"experience,omitempty"`
	Specialization string        `gorm:"size:255" json:"specialization,omitempty"`
	Status         TeacherStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	IsActive       bool          `gorm:"default:true" json:"is_active"`

	// Relationships
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
