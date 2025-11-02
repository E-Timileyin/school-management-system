package model

import "time"

type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "present"
	AttendanceStatusAbsent  AttendanceStatus = "absent"
	AttendanceStatusLate    AttendanceStatus = "late"
	AttendanceStatusExcused AttendanceStatus = "excused"
)

type Attendance struct {
	Base
	StudentID     uint            `gorm:"not null" json:"student_id"`
	ClassID       uint            `gorm:"not null" json:"class_id"`
	SectionID     uint            `gorm:"not null" json:"section_id"`
	SubjectID     *uint           `gorm:"index" json:"subject_id,omitempty"` // Optional, for subject-specific attendance
	Date          time.Time       `gorm:"type:date;not null" json:"date"`
	Status        AttendanceStatus `gorm:"type:varchar(20);default:'present'" json:"status"`
	Remarks       string          `gorm:"type:text" json:"remarks,omitempty"`
	MarkedBy      uint            `gorm:"not null" json:"marked_by"` // UserID of the staff who marked attendance
	AcademicYearID uint           `gorm:"not null" json:"academic_year_id"`

	// Relationships
	Student      *Student      `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	Class        *Class        `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Section      *Section      `gorm:"foreignKey:SectionID" json:"section,omitempty"`
	Subject      *Subject      `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	MarkedByUser *User         `gorm:"foreignKey:MarkedBy" json:"marked_by_user,omitempty"`
	AcademicYear *AcademicYear `gorm:"foreignKey:AcademicYearID" json:"academic_year,omitempty"`
}
