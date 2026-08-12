package handler

import (
	"crypto/subtle"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yigeyingshi/yige-admin-server/internal/auth"
	"github.com/yigeyingshi/yige-admin-server/internal/config"
)

type AuthHandler struct {
	cfg    *config.Config
	tokens *auth.TokenService
}

func NewAuthHandler(cfg *config.Config, tokens *auth.TokenService) *AuthHandler {
	return &AuthHandler{cfg: cfg, tokens: tokens}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required,max=64"`
		Password string `json:"password" binding:"required,max=256"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		Error(c, http.StatusBadRequest, "INVALID_INPUT", "请输入账号和密码")
		return
	}
	userOK := subtle.ConstantTimeCompare([]byte(input.Username), []byte(h.cfg.Username)) == 1
	passwordOK := subtle.ConstantTimeCompare([]byte(input.Password), []byte(h.cfg.Password)) == 1
	if !userOK || !passwordOK {
		time.Sleep(120 * time.Millisecond)
		Error(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "账号或密码错误")
		return
	}
	token, err := h.tokens.Issue(h.cfg.Username)
	if err != nil {
		Error(c, http.StatusInternalServerError, "INTERNAL_ERROR", "暂时无法登录")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": token, "expiresIn": int(h.cfg.TokenTTL.Seconds()), "user": gin.H{"username": h.cfg.Username, "displayName": "内容管理员"}}})
}
