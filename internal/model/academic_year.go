package model

import "time"

type AcademicYearStatus string

const (
	AcademicYearStatusUpcoming AcademicYearStatus = "upcoming"
	AcademicYearStatusActive   AcademicYearStatus = "active"
	AcademicYearStatusCompleted AcademicYearStatus = "completed"
)

type AcademicYear struct {
	Base
	Name        string            `gorm:"size:50;not null;uniqueIndex" json:"name"` // e.g., "2023-2024"
	StartDate   time.Time         `gorm:"not null" json:"start_date"`
	EndDate     time.Time         `gorm:"not null" json:"end_date"`
	Status      AcademicYearStatus `gorm:"type:varchar(20);default:'upcoming'" json:"status"`
	IsCurrent   bool              `gorm:"default:false" json:"is_current"`
	Description string            `gorm:"type:text" json:"description,omitempty"`
}
