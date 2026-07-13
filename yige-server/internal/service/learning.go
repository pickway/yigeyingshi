package service

import (
	"github.com/yigeyingshi/yige-server/internal/model"
	"gorm.io/gorm"
)

type LearningService struct {
	db *gorm.DB
}

func NewLearningService(db *gorm.DB) *LearningService {
	return &LearningService{db: db}
}

func (s *LearningService) ListCourses(category, level string) ([]model.LearningCourse, error) {
	var list []model.LearningCourse
	q := s.db.Model(&model.LearningCourse{}).Where("is_active = ?", true)
	if category != "" && category != "全部" {
		q = q.Where("category = ?", category)
	}
	if level != "" {
		q = q.Where("level = ?", level)
	}
	err := q.Order("id ASC").Find(&list).Error
	return list, err
}

func (s *LearningService) ListPaths() ([]model.LearningPath, error) {
	var list []model.LearningPath
	err := s.db.Find(&list).Error
	return list, err
}
