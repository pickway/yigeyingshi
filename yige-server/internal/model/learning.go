package model

import (
	"time"
)

// LearningCourse 学习课程/资源
type LearningCourse struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:128;not null" json:"title"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Category    string    `gorm:"size:32;index" json:"category"`
	Level       string    `gorm:"size:16" json:"level"`
	Duration    string    `gorm:"size:32" json:"duration,omitempty"`
	Lessons     int       `gorm:"default:0" json:"lessons"`
	Icon        string    `gorm:"size:32" json:"icon,omitempty"`
	IsActive    bool      `gorm:"default:true" json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// LearningPath 学习路径
type LearningPath struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:128;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CourseIDs   string    `gorm:"type:text" json:"courseIds,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
