package model

type Section struct {
	Base
	Name           string `gorm:"size:10;not null" json:"name"` // A, B, C...
	ClassID        uint   `gorm:"not null" json:"class_id"`
	ClassTeacherID *uint  `gorm:"index" json:"class_teacher_id,omitempty"`
	Capacity      int    `gorm:"default:40" json:"capacity"`
	IsActive      bool   `gorm:"default:true" json:"is_active"`

	// Relationships
	Class        *Class  `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	ClassTeacher *Teacher `gorm:"foreignKey:ClassTeacherID" json:"class_teacher,omitempty"`
}
