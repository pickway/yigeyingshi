package model

import (
	"time"
)

// Article 文章（影评/学习笔记/AI 文章）
type Article struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:256;not null" json:"title"`
	Category    string    `gorm:"size:32;not null;index" json:"category"`
	Tag         string    `gorm:"size:32" json:"tag,omitempty"`
	Summary     string    `gorm:"type:text" json:"summary,omitempty"`
	Content     string    `gorm:"type:text" json:"content,omitempty"`
	CoverColor  string    `gorm:"size:32" json:"coverColor,omitempty"`
	Author      string    `gorm:"size:64" json:"author,omitempty"`
	ViewCount   int       `gorm:"default:0" json:"viewCount"`
	PublishedAt time.Time `json:"publishedAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Category 枚举: movie-review, learning-note, ai-article
