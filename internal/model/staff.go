package model

import "time"

type StaffType string

type StaffStatus string

const (
	StaffTypeAdministrative StaffType = "administrative"
	StaffTypeSupport       StaffType = "support"
	StaffTypeLibrarian     StaffType = "librarian"
	StaffTypeOther         StaffType = "other"

	StaffStatusActive   StaffStatus = "active"
	StaffStatusInactive StaffStatus = "inactive"
	StaffStatusOnLeave  StaffStatus = "on_leave"
)

type Staff struct {
	Base
	UserID         uint        `gorm:"not null;uniqueIndex" json:"user_id"`
	EmployeeID     string      `gorm:"size:50;unique;not null" json:"employee_id"`
	StaffType      StaffType   `gorm:"type:varchar(30);not null" json:"staff_type"`
	Department     string      `gorm:"size:100" json:"department,omitempty"`
	Designation    string      `gorm:"size:100;not null" json:"designation"`
	JoiningDate    time.Time   `gorm:"not null" json:"joining_date"`
	Qualification  string      `gorm:"size:255" json:"qualification,omitempty"`
	Experience     string      `gorm:"size:100" json:"experience,omitempty"`
	Status         StaffStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	IsActive       bool        `gorm:"default:true" json:"is_active"`

	// Relationships
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
