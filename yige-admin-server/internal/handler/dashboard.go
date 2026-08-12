package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-admin-server/internal/model"
	"gorm.io/gorm"
)

type DashboardHandler struct{ db *gorm.DB }

func NewDashboardHandler(db *gorm.DB) *DashboardHandler { return &DashboardHandler{db: db} }

func (h *DashboardHandler) Get(c *gin.Context) {
	counts := map[string]int64{}
	models := map[string]interface{}{"movies": &model.Movie{}, "articles": &model.Article{}, "courses": &model.LearningCourse{}, "paths": &model.LearningPath{}, "tools": &model.AiTool{}, "subscribers": &model.Newsletter{}}
	for key, item := range models {
		var count int64
		if err := h.db.Model(item).Count(&count).Error; err != nil {
			Error(c, http.StatusInternalServerError, "DATABASE_ERROR", "暂时无法加载统计")
			return
		}
		counts[key] = count
	}
	var recentArticles []model.Article
	if err := h.db.Order("updated_at DESC").Limit(5).Find(&recentArticles).Error; err != nil {
		Error(c, http.StatusInternalServerError, "DATABASE_ERROR", "暂时无法加载最近文章")
		return
	}
	var recentMovies []model.Movie
	if err := h.db.Order("updated_at DESC").Limit(5).Find(&recentMovies).Error; err != nil {
		Error(c, http.StatusInternalServerError, "DATABASE_ERROR", "暂时无法加载最近影片")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"counts": counts, "recentArticles": recentArticles, "recentMovies": recentMovies}})
}
