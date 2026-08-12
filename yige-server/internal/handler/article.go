package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type ArticleHandler struct {
	service *service.ArticleService
}

func (h *ArticleHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		writeError(c, http.StatusBadRequest, "INVALID_ID", "文章编号无效")
		return
	}
	article, err := h.service.GetByID(uint(id))
	if err != nil {
		writeError(c, http.StatusNotFound, "NOT_FOUND", "没有找到这篇文章")
		return
	}
	related, err := h.service.Related(article, 3)
	if err != nil {
		writeError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "暂时无法加载文章")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": article, "related": related})
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
