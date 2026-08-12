package service

import (
	"strings"

	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

type SearchResults struct {
	Movies   []model.Movie          `json:"movies"`
	Articles []model.Article        `json:"articles"`
	Courses  []model.LearningCourse `json:"courses"`
	Tools    []model.AiTool         `json:"tools"`
}

type SearchService struct {
	db *gorm.DB
}

func NewSearchService(db *gorm.DB) *SearchService {
	return &SearchService{db: db}
}

func (s *SearchService) Search(keyword string, limit int) (*SearchResults, error) {
	pattern := "%" + strings.TrimSpace(keyword) + "%"
	results := &SearchResults{
		Movies:   []model.Movie{},
		Articles: []model.Article{},
		Courses:  []model.LearningCourse{},
		Tools:    []model.AiTool{},
	}
	queries := []struct {
		target interface{}
		query  string
		args   []interface{}
		order  string
	}{
		{&results.Movies, "title LIKE ? OR original_title LIKE ? OR director LIKE ? OR genres LIKE ?", []interface{}{pattern, pattern, pattern, pattern}, "rating DESC, year DESC"},
		{&results.Articles, "title LIKE ? OR summary LIKE ? OR tag LIKE ?", []interface{}{pattern, pattern, pattern}, "published_at DESC"},
		{&results.Courses, "is_active = ? AND (title LIKE ? OR description LIKE ? OR category LIKE ?)", []interface{}{true, pattern, pattern, pattern}, "id ASC"},
		{&results.Tools, "name LIKE ? OR description LIKE ? OR tags LIKE ?", []interface{}{pattern, pattern, pattern}, "featured DESC, id ASC"},
	}
	for _, query := range queries {
		if err := s.db.Where(query.query, query.args...).Order(query.order).Limit(limit).Find(query.target).Error; err != nil {
			return nil, err
		}
	}
	return results, nil
}
