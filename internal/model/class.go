package model

type Class struct {
	Base
	Name          string `gorm:"size:50;not null;uniqueIndex" json:"name"` // e.g., "Class 1", "Class 2"
	NumericValue  int    `gorm:"not null;uniqueIndex" json:"numeric_value"` // 1, 2, 3...
	Description   string `gorm:"type:text" json:"description,omitempty"`
	IsActive      bool   `gorm:"default:true" json:"is_active"`
	
	// Relationships
	Sections []Section `gorm:"foreignKey:ClassID" json:"sections,omitempty"`
}
