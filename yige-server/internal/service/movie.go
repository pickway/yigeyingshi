package service

import (
	"strings"

	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) List(genre, keyword, sort string, page, pageSize int) ([]model.Movie, int64, error) {
	var list []model.Movie
	var total int64

	q := s.db.Model(&model.Movie{})
	if genre != "" && genre != "全部" {
		q = q.Where("genres LIKE ?", "%"+genre+"%")
	}
	if keyword != "" {
		pattern := "%" + strings.TrimSpace(keyword) + "%"
		q = q.Where(
			"title LIKE ? OR original_title LIKE ? OR director LIKE ? OR description LIKE ? OR genres LIKE ?",
			pattern, pattern, pattern, pattern, pattern,
		)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	switch sort {
	case "year_desc":
		q = q.Order("year DESC, rating DESC")
	case "year_asc":
		q = q.Order("year ASC, rating DESC")
	case "popular":
		q = q.Order("rating DESC, year DESC, id ASC")
	default:
		q = q.Order("rating DESC, year DESC, id ASC")
	}

	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	err := q.Offset(offset).Limit(pageSize).Find(&list).Error
	return list, total, err
}

func (s *MovieService) Featured(limit int) ([]model.Movie, error) {
	var list []model.Movie
	err := s.db.Where("is_featured = ?", true).
		Order("rating DESC").
		Limit(limit).
		Find(&list).Error
	return list, err
}

func (s *MovieService) GetByID(id uint) (*model.Movie, error) {
	var m model.Movie
	err := s.db.First(&m, id).Error
	return &m, err
}
