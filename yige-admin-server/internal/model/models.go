package model

import "time"

type Movie struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	Title         string    `gorm:"size:128;not null" json:"title" binding:"required,max=128"`
	OriginalTitle string    `gorm:"size:128" json:"originalTitle"`
	Year          int       `json:"year"`
	Genres        string    `gorm:"size:128" json:"genres"`
	Rating        float64   `json:"rating"`
	Description   string    `gorm:"type:text" json:"description"`
	PosterColor   string    `gorm:"size:32" json:"posterColor"`
	PosterIcon    string    `gorm:"size:32" json:"posterIcon"`
	Duration      int       `json:"duration"`
	Director      string    `gorm:"size:64" json:"director"`
	IsFeatured    bool      `json:"isFeatured"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type Article struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:256;not null" json:"title" binding:"required,max=256"`
	Category    string    `gorm:"size:32;not null;index" json:"category" binding:"required"`
	Tag         string    `gorm:"size:32" json:"tag"`
	Summary     string    `gorm:"type:text" json:"summary"`
	Content     string    `gorm:"type:text" json:"content"`
	CoverColor  string    `gorm:"size:32" json:"coverColor"`
	Author      string    `gorm:"size:64" json:"author"`
	ViewCount   int       `json:"viewCount"`
	PublishedAt time.Time `json:"publishedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type LearningCourse struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:128;not null" json:"title" binding:"required,max=128"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"size:32;index" json:"category"`
	Level       string    `gorm:"size:16" json:"level"`
	Duration    string    `gorm:"size:32" json:"duration"`
	Lessons     int       `json:"lessons"`
	Icon        string    `gorm:"size:32" json:"icon"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (LearningCourse) TableName() string { return "learning_courses" }

type LearningPath struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:128;not null" json:"name" binding:"required,max=128"`
	Description string    `gorm:"type:text" json:"description"`
	CourseIDs   string    `gorm:"type:text" json:"courseIds"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (LearningPath) TableName() string { return "learning_paths" }

type AiTool struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:64;not null" json:"name" binding:"required,max=64"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"size:32;index" json:"category"`
	Icon        string    `gorm:"size:32" json:"icon"`
	Tags        string    `gorm:"size:128" json:"tags"`
	Url         string    `gorm:"size:256" json:"url"`
	IsFree      bool      `json:"isFree"`
	Featured    bool      `json:"featured"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (AiTool) TableName() string { return "ai_tools" }

type Newsletter struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Email     string    `gorm:"size:256;uniqueIndex;not null" json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

func (Newsletter) TableName() string { return "newsletters" }
