package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type AiHandler struct {
	toolService   *service.AiToolService
	newsletterSvc *service.NewsletterService
}

func NewAiHandler(tool *service.AiToolService, newsletter *service.NewsletterService) *AiHandler {
	return &AiHandler{toolService: tool, newsletterSvc: newsletter}
}

func (h *AiHandler) ListTools(c *gin.Context) {
	category := c.Query("category")
	list, err := h.toolService.List(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *AiHandler) FeaturedTools(c *gin.Context) {
	list, err := h.toolService.Featured(3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

type newsletterReq struct {
	Email string `json:"email" binding:"required,email,max=254"`
}

func (h *AiHandler) Subscribe(c *gin.Context) {
	var req newsletterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, http.StatusBadRequest, "INVALID_EMAIL", "请输入有效的邮箱地址")
		return
	}
	if err := h.newsletterSvc.Subscribe(strings.TrimSpace(req.Email)); err != nil {
		writeError(c, http.StatusConflict, "ALREADY_SUBSCRIBED", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "订阅成功"})
}
