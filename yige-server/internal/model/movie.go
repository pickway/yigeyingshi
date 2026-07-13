package model

import (
	"time"
)

// Movie 影视资源
type Movie struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:128;not null" json:"title"`
	OriginalTitle string  `gorm:"size:128" json:"originalTitle,omitempty"`
	Year        int       `json:"year"`
	Genres      string    `gorm:"size:128" json:"genres"`
	Rating      float64   `gorm:"type:decimal(3,1)" json:"rating"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	PosterColor string    `gorm:"size:32" json:"posterColor,omitempty"`
	PosterIcon  string    `gorm:"size:32" json:"posterIcon,omitempty"`
	Duration    int       `json:"duration,omitempty"`
	Director    string    `gorm:"size:64" json:"director,omitempty"`
	IsFeatured  bool      `gorm:"default:false" json:"isFeatured"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
