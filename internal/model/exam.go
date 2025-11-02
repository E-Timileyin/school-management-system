package model

import "time"

type ExamType string

const (
	ExamTypeQuarterly  ExamType = "quarterly"
	ExamTypeHalfYearly ExamType = "half_yearly"
	ExamTypeAnnual     ExamType = "annual"
	ExamTypeUnitTest   ExamType = "unit_test"
	ExamTypeMidTerm    ExamType = "mid_term"
)

type ExamStatus string

const (
	ExamStatusDraft     ExamStatus = "draft"
	ExamStatusScheduled ExamStatus = "scheduled"
	ExamStatusOngoing   ExamStatus = "ongoing"
	ExamStatusCompleted ExamStatus = "completed"
	ExamStatusCancelled ExamStatus = "cancelled"
)

type Exam struct {
	Base
	Name           string     `gorm:"size:100;not null" json:"name"`
	ExamType       ExamType   `gorm:"type:varchar(20);not null" json:"exam_type"`
	StartDate      time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate        time.Time  `gorm:"type:date;not null" json:"end_date"`
	AcademicYearID uint       `gorm:"not null" json:"academic_year_id"`
	Status         ExamStatus `gorm:"type:varchar(20);default:'draft'" json:"status"`
	Description    string     `gorm:"type:text" json:"description,omitempty"`
	IsPublished    bool       `gorm:"default:false" json:"is_published"`

	// Relationships
	AcademicYear *AcademicYear `gorm:"foreignKey:AcademicYearID" json:"academic_year,omitempty"`
	ExamSubjects []ExamSubject `gorm:"foreignKey:ExamID" json:"exam_subjects,omitempty"`
}

type ExamSubject struct {
	Base
	ExamID         uint      `gorm:"not null;index" json:"exam_id"`
	SubjectID      uint      `gorm:"not null" json:"subject_id"`
	ClassID        uint      `gorm:"not null" json:"class_id"`
	ExamDate       time.Time `gorm:"type:date;not null" json:"exam_date"`
	StartTime      string    `gorm:"size:10;not null" json:"start_time"` // Format: "HH:MM"
	EndTime        string    `gorm:"size:10;not null" json:"end_time"`   // Format: "HH:MM"
	MaxMarks       float64   `gorm:"not null;default:100" json:"max_marks"`
	PassingMarks   float64   `gorm:"not null;default:35" json:"passing_marks"`
	RoomNumber     string    `gorm:"size:20" json:"room_number,omitempty"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`

	// Relationships
	Exam          *Exam          `gorm:"foreignKey:ExamID" json:"exam,omitempty"`
	Subject       *Subject       `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	Class         *Class         `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	ExamResults   []ExamResult   `gorm:"foreignKey:ExamSubjectID" json:"exam_results,omitempty"`
}

type ExamResult struct {
	Base
	ExamSubjectID uint    `gorm:"not null;index" json:"exam_subject_id"`
	StudentID     uint    `gorm:"not null;index" json:"student_id"`
	MarksObtained float64 `gorm:"not null;default:0" json:"marks_obtained"`
	Grade         string  `gorm:"size:5" json:"grade,omitempty"`
	Remarks       string  `gorm:"type:text" json:"remarks,omitempty"`
	IsPublished   bool    `gorm:"default:false" json:"is_published"`
	PublishedAt   *time.Time `gorm:"default:null" json:"published_at,omitempty"`

	// Relationships
	ExamSubject *ExamSubject `gorm:"foreignKey:ExamSubjectID" json:"exam_subject,omitempty"`
	Student     *Student     `gorm:"foreignKey:StudentID" json:"student,omitempty"`
}
