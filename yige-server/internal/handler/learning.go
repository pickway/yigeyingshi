package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type LearningHandler struct {
	service *service.LearningService
}

func NewLearningHandler(s *service.LearningService) *LearningHandler {
	return &LearningHandler{service: s}
}

func (h *LearningHandler) ListCourses(c *gin.Context) {
	category := c.Query("category")
	level := c.Query("level")
	list, err := h.service.ListCourses(category, level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *LearningHandler) ListPaths(c *gin.Context) {
	list, err := h.service.ListPaths()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}
