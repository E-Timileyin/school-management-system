package model

type DayOfWeek int

const (
	Sunday DayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Timetable struct {
	Base
	ClassID       uint     `gorm:"not null" json:"class_id"`
	SectionID     uint     `gorm:"not null" json:"section_id"`
	SubjectID     uint     `gorm:"not null" json:"subject_id"`
	TeacherID     uint     `gorm:"not null" json:"teacher_id"`
	DayOfWeek     DayOfWeek `gorm:"not null" json:"day_of_week"` // 0-6 (Sunday-Saturday)
	PeriodNumber  int      `gorm:"not null" json:"period_number"`
	StartTime     string   `gorm:"size:10;not null" json:"start_time"` // Format: "HH:MM"
	EndTime       string   `gorm:"size:10;not null" json:"end_time"`   // Format: "HH:MM"
	AcademicYearID uint     `gorm:"not null" json:"academic_year_id"`
	IsActive      bool     `gorm:"default:true" json:"is_active"`

	// Relationships
	Class        *Class        `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Section      *Section      `gorm:"foreignKey:SectionID" json:"section,omitempty"`
	Subject      *Subject      `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	Teacher      *Teacher      `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	AcademicYear *AcademicYear `gorm:"foreignKey:AcademicYearID" json:"academic_year,omitempty"`
}
