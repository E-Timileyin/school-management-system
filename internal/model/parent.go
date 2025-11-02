package model

type ParentType string

const (
	ParentTypeFather   ParentType = "father"
	ParentTypeMother   ParentType = "mother"
	ParentTypeGuardian ParentType = "guardian"
)

type Parent struct {
	Base
	UserID     uint       `gorm:"not null;uniqueIndex" json:"user_id"`
	Type       ParentType `gorm:"type:varchar(20);not null" json:"type"`
	Occupation string     `gorm:"size:100" json:"occupation,omitempty"`
	IsPrimary  bool       `gorm:"default:false" json:"is_primary"`

	// Relationships
	User     *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Students []*Student `gorm:"many2many:student_parents;" json:"students,omitempty"`
}
