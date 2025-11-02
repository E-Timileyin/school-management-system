package model

type SubjectType string

const (
	SubjectTypeCore           SubjectType = "core"
	SubjectTypeElective       SubjectType = "elective"
	SubjectTypeExtraCurricular SubjectType = "extra_curricular"
)

type Subject struct {
	Base
	Name        string      `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Code        string      `gorm:"size:20;unique;not null" json:"code"`
	Type        SubjectType `gorm:"type:varchar(20);default:'core'" json:"type"`
	Description string      `gorm:"type:text" json:"description,omitempty"`
	IsActive    bool        `gorm:"default:true" json:"is_active"`

	// Relationships
	Classes []*Class `gorm:"many2many:class_subjects;" json:"classes,omitempty"`
}

type ClassSubject struct {
	Base
	ClassID       uint `gorm:"not null;uniqueIndex:idx_class_subject_teacher" json:"class_id"`
	SubjectID     uint `gorm:"not null;uniqueIndex:idx_class_subject_teacher" json:"subject_id"`
	TeacherID     uint `gorm:"not null;uniqueIndex:idx_class_subject_teacher" json:"teacher_id"`
	AcademicYearID uint `gorm:"not null" json:"academic_year_id"`
	IsActive      bool `gorm:"default:true" json:"is_active"`

	// Relationships
	Class        *Class        `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Subject      *Subject      `gorm:"foreignKey:SubjectID" json:"subject,omitempty"`
	Teacher      *Teacher      `gorm:"foreignKey:TeacherID" json:"teacher,omitempty"`
	AcademicYear *AcademicYear `gorm:"foreignKey:AcademicYearID" json:"academic_year,omitempty"`
}
