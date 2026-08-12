package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(s *service.MovieService) *MovieHandler {
	return &MovieHandler{service: s}
}

func (h *MovieHandler) List(c *gin.Context) {
	genre := c.Query("genre")
	keyword := c.Query("keyword")
	sort := c.DefaultQuery("sort", "rating_desc")
	page, pageErr := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, pageSizeErr := strconv.Atoi(c.DefaultQuery("pageSize", "12"))
	if pageErr != nil || pageSizeErr != nil || page < 1 || pageSize < 1 || pageSize > 48 {
		writeError(c, http.StatusBadRequest, "INVALID_PAGINATION", "分页参数无效")
		return
	}
	if len([]rune(keyword)) > 100 {
		writeError(c, http.StatusBadRequest, "INVALID_KEYWORD", "搜索关键词过长")
		return
	}
	validSorts := map[string]bool{"rating_desc": true, "year_desc": true, "year_asc": true, "popular": true}
	if !validSorts[sort] {
		writeError(c, http.StatusBadRequest, "INVALID_SORT", "排序方式无效")
		return
	}

	list, total, err := h.service.List(genre, keyword, sort, page, pageSize)
	if err != nil {
		writeError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "暂时无法加载影片")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":     list,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

func (h *MovieHandler) Featured(c *gin.Context) {
	list, err := h.service.Featured(3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *MovieHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		writeError(c, http.StatusBadRequest, "INVALID_ID", "影片编号无效")
		return
	}
	movie, err := h.service.GetByID(uint(id))
	if err != nil {
		writeError(c, http.StatusNotFound, "NOT_FOUND", "没有找到这部影片")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": movie})
}
