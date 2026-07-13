package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type ArticleHandler struct {
	service *service.ArticleService
}

func NewArticleHandler(s *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: s}
}

func (h *ArticleHandler) List(c *gin.Context) {
	category := c.Query("category")
	tag := c.Query("tag")

	list, err := h.service.List(category, tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *ArticleHandler) Latest(c *gin.Context) {
	limit := 2
	if l := c.Query("limit"); l != "" {
		if n, err := parseInt(l); err == nil && n > 0 {
			limit = n
		}
	}
	list, err := h.service.Latest(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func parseInt(s string) (int, error) {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, nil
		}
		n = n*10 + int(c-'0')
	}
	return n, nil
}
