package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type SearchHandler struct {
	service *service.SearchService
}

func NewSearchHandler(searchService *service.SearchService) *SearchHandler {
	return &SearchHandler{service: searchService}
}

func (h *SearchHandler) Search(c *gin.Context) {
	keyword := strings.TrimSpace(c.Query("q"))
	if len([]rune(keyword)) < 2 || len([]rune(keyword)) > 100 {
		writeError(c, http.StatusBadRequest, "INVALID_QUERY", "请输入 2 到 100 个字符")
		return
	}
	results, err := h.service.Search(keyword, 5)
	if err != nil {
		writeError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "暂时无法搜索")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": results, "query": keyword})
}
