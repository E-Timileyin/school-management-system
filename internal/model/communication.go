package model

import "time"

type CommunicationType string

type AudienceType string

const (
	CommunicationTypeNotice CommunicationType = "notice"
	CommunicationTypeEvent  CommunicationType = "event"
	CommunicationTypeNews   CommunicationType = "news"

	AudienceAll      AudienceType = "all"
	AudienceStudents AudienceType = "students"
	AudienceTeachers AudienceType = "teachers"
	AudienceParents  AudienceType = "parents"
	AudienceStaff    AudienceType = "staff"
)

type Communication struct {
	Base
	Title        string           `gorm:"size:255;not null" json:"title"`
	Content      string           `gorm:"type:text;not null" json:"content"`
	CommType     CommunicationType `gorm:"type:varchar(20);not null" json:"comm_type"`
	Audience     AudienceType     `gorm:"type:varchar(20);not null" json:"audience"`
	StartDate    *time.Time       `gorm:"type:timestamp" json:"start_date,omitempty"`
	EndDate      *time.Time       `gorm:"type:timestamp" json:"end_date,omitempty"`
	IsPublished  bool             `gorm:"default:false" json:"is_published"`
	PublishedAt  *time.Time       `gorm:"type:timestamp" json:"published_at,omitempty"`
	AuthorID     uint             `gorm:"not null" json:"author_id"`
	TargetClassID *uint           `gorm:"index" json:"target_class_id,omitempty"`
	TargetUserID  *uint           `gorm:"index" json:"target_user_id,omitempty"`

	// For events
	Location     string           `gorm:"size:255" json:"location,omitempty"`
	IsAllDay     bool             `gorm:"default:false" json:"is_all_day"`

	// Relationships
	Author      *User  `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	TargetClass *Class `gorm:"foreignKey:TargetClassID" json:"target_class,omitempty"`
	Attachments []CommunicationAttachment `gorm:"foreignKey:CommunicationID" json:"attachments,omitempty"`
}

type CommunicationAttachment struct {
	Base
	CommunicationID uint   `gorm:"not null;index" json:"communication_id"`
	FileName       string `gorm:"size:255;not null" json:"file_name"`
	FileURL        string `gorm:"type:text;not null" json:"file_url"`
	FileType       string `gorm:"size:100" json:"file_type,omitempty"`
	FileSize       int64  `gorm:"default:0" json:"file_size"`
}
