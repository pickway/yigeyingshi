package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-server/internal/service"
)

type AiHandler struct {
	toolService  *service.AiToolService
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
	Email string `json:"email" binding:"required,email"`
}

func (h *AiHandler) Subscribe(c *gin.Context) {
	var req newsletterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}
	if err := h.newsletterSvc.Subscribe(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "订阅成功"})
}
