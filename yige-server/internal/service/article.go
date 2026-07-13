package service

import (
	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

func (s *ArticleService) List(category, tag string) ([]model.Article, error) {
	var list []model.Article
	q := s.db.Model(&model.Article{})
	if category != "" {
		q = q.Where("category = ?", category)
	}
	if tag != "" {
		q = q.Where("tag = ?", tag)
	}
	err := q.Order("published_at DESC").Find(&list).Error
	return list, err
}

func (s *ArticleService) Latest(limit int) ([]model.Article, error) {
	var list []model.Article
	err := s.db.Order("published_at DESC").Limit(limit).Find(&list).Error
	return list, err
}
