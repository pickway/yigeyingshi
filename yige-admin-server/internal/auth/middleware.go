package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(service *TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			abort(c, "AUTH_REQUIRED", "请先登录")
			return
		}
		claims, err := service.Verify(strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			abort(c, "INVALID_TOKEN", "登录已过期，请重新登录")
			return
		}
		c.Set("admin", claims.Subject)
		c.Next()
	}
}

func abort(c *gin.Context, code, message string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": gin.H{"code": code, "message": message}})
}
