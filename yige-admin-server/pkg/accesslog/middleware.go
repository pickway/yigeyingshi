package accesslog

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Record 一条 HTTP 请求日志，字段顺序与 accesslog.go Write 中的 zap.Field 一致
type Record struct {
	Service   string  `json:"service"`
	Method    string  `json:"method"`
	Path      string  `json:"path"`
	Query     string  `json:"query"`
	Status    int     `json:"status"`
	LatencyMs float64 `json:"latency_ms"`
	ClientIP  string  `json:"client_ip"`
	UserAgent string  `json:"user_agent"`
	RequestID string  `json:"request_id"`
	BytesOut  int     `json:"bytes_out"`
	Error     string  `json:"error,omitempty"`
}

// Middleware 返回 gin 中间件，每次请求结束后异步写一条 JSON 行日志到本地文件
// serviceName 用于区分来源（yige-server / yige-admin-server）
func Middleware(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 透传 X-Request-ID（如果上游带）
		requestID := c.GetHeader("X-Request-ID")
		if requestID != "" {
			c.Header("X-Request-ID", requestID)
		}

		c.Next()

		latencyMs := float64(time.Since(start).Nanoseconds()) / 1e6
		bytesOut := c.Writer.Size()
		if bytesOut < 0 {
			bytesOut = 0
		}

		rec := Record{
			Service:   serviceName,
			Method:    c.Request.Method,
			Path:      path,
			Query:     query,
			Status:    c.Writer.Status(),
			LatencyMs: latencyMs,
			ClientIP:  c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			RequestID: requestID,
			BytesOut:  bytesOut,
		}
		if len(c.Errors) > 0 {
			rec.Error = c.Errors.String()
		}

		// zap 本身是并发安全且带 buffer 的，直接同步写即可
		if l := Default(); l != nil {
			l.Write(rec)
		}
	}
}
