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

func (s *MovieService) List(genre, sort string, page, pageSize int) ([]model.Movie, int64, error) {
	var list []model.Movie
	var total int64

	q := s.db.Model(&model.Movie{})
	if genre != "" && genre != "全部" {
		q = q.Where("genres LIKE ?", "%"+genre+"%")
	}

	q.Count(&total)

	switch sort {
	case "year_desc":
		q = q.Order("year DESC, rating DESC")
	case "year_asc":
		q = q.Order("year ASC, rating DESC")
	default:
		q = q.Order("rating DESC, year DESC")
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

func (s *MovieService) Create(m *model.Movie) error {
	return s.db.Create(m).Error
}

func splitGenres(genres string) []string {
	trimmed := strings.TrimSpace(genres)
	if trimmed == "" {
		return nil
	}
	parts := strings.Split(trimmed, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
