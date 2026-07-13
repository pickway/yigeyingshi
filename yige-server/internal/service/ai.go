package service

import (
	"errors"

	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

type AiToolService struct {
	db *gorm.DB
}

func NewAiToolService(db *gorm.DB) *AiToolService {
	return &AiToolService{db: db}
}

func (s *AiToolService) List(category string) ([]model.AiTool, error) {
	var list []model.AiTool
	q := s.db.Model(&model.AiTool{})
	if category != "" {
		q = q.Where("category = ?", category)
	}
	err := q.Order("id ASC").Find(&list).Error
	return list, err
}

func (s *AiToolService) Featured(limit int) ([]model.AiTool, error) {
	var list []model.AiTool
	err := s.db.Where("featured = ?", true).Limit(limit).Find(&list).Error
	return list, err
}

type NewsletterService struct {
	db *gorm.DB
}

func NewNewsletterService(db *gorm.DB) *NewsletterService {
	return &NewsletterService{db: db}
}

func (s *NewsletterService) Subscribe(email string) error {
	var existing model.Newsletter
	err := s.db.Where("email = ?", email).First(&existing).Error
	if err == nil {
		return errors.New("该邮箱已订阅")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return s.db.Create(&model.Newsletter{Email: email, Active: true}).Error
}
