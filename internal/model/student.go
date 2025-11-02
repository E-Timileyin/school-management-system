package model

import "time"

type StudentStatus string

const (
	StudentStatusActive     StudentStatus = "active"
	StudentStatusInactive   StudentStatus = "inactive"
	StudentStatusGraduated  StudentStatus = "graduated"
	StudentStatusTransferred StudentStatus = "transferred"
)

type Student struct {
	Base
	UserID        uint          `gorm:"not null;uniqueIndex" json:"user_id"`
	AdmissionNo   string        `gorm:"size:50;unique;not null" json:"admission_no"`
	AdmissionDate time.Time     `gorm:"not null" json:"admission_date"`
	ClassID       uint          `gorm:"not null" json:"class_id"`
	SectionID     uint          `gorm:"not null" json:"section_id"`
	RollNumber    int           `gorm:"not null" json:"roll_number"`
	Status        StudentStatus `gorm:"type:varchar(20);default:'active'" json:"status"`
	IsActive      bool          `gorm:"default:true" json:"is_active"`

	// Relationships
	User    *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Class   *Class    `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Section *Section  `gorm:"foreignKey:SectionID" json:"section,omitempty"`
	Parents []*Parent `gorm:"many2many:student_parents;" json:"parents,omitempty"`
}
