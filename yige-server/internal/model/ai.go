package model

import (
	"time"
)

// AiTool AI 工具推荐
type AiTool struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:64;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Category    string    `gorm:"size:32;index" json:"category"`
	Icon        string    `gorm:"size:32" json:"icon,omitempty"`
	Tags        string    `gorm:"size:128" json:"tags,omitempty"`
	Url         string    `gorm:"size:256" json:"url,omitempty"`
	IsFree      bool      `gorm:"default:true" json:"isFree"`
	Featured    bool      `gorm:"default:false" json:"featured"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Newsletter 订阅者
type Newsletter struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Email     string    `gorm:"size:256;uniqueIndex;not null" json:"email"`
	Active    bool      `gorm:"default:true" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}
