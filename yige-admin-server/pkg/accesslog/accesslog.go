// Package accesslog 用 zap + lumberjack 把 HTTP 请求日志按 JSON 行写到本地文件，
// 供 Filebeat 采集到 Elasticsearch。字段名遵循 ECS（Elastic Common Schema）规范，
// 与索引模板 yige-request-logs 的字段定义对齐。
package accesslog

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config 访问日志配置
type Config struct {
	// FilePath 日志文件路径，例如 /var/log/yige/yige-admin-server-access.log 或 logs/access.log
	FilePath string
	// MaxSizeMB 单个日志文件最大体积（MB），超过后滚动切分
	MaxSizeMB int
	// MaxBackups 保留的旧日志文件数量
	MaxBackups int
	// MaxAgeDays 旧日志保留天数
	MaxAgeDays int
	// Compress 是否 gzip 压缩归档日志
	Compress bool
}

// Logger 访问日志记录器
type Logger struct {
	z  *zap.Logger
	mu sync.Mutex
}

var (
	defaultLogger *Logger
	initOnce      sync.Once
	initErr       error
)

// New 创建一个新的 Logger，并确保父目录存在
func New(cfg Config) (*Logger, error) {
	if cfg.FilePath == "" {
		return nil, fmt.Errorf("accesslog: FilePath 不能为空")
	}
	if err := os.MkdirAll(filepath.Dir(cfg.FilePath), 0o755); err != nil {
		return nil, fmt.Errorf("accesslog: 创建日志目录失败: %w", err)
	}

	maxSize := cfg.MaxSizeMB
	if maxSize == 0 {
		maxSize = 100
	}
	maxBackups := cfg.MaxBackups
	if maxBackups == 0 {
		maxBackups = 5
	}
	maxAge := cfg.MaxAgeDays
	if maxAge == 0 {
		maxAge = 7
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   cfg.Compress,
		LocalTime:  true,
	})

	// ECS 规范字段名：service.name / client.ip / user_agent.original
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "@timestamp"
	encoderCfg.LevelKey = ""
	encoderCfg.NameKey = ""
	encoderCfg.CallerKey = ""
	encoderCfg.MessageKey = ""
	encoderCfg.StacktraceKey = ""
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		w,
		zap.InfoLevel,
	)
	return &Logger{z: zap.New(core)}, nil
}

// Init 初始化全局默认 Logger（幂等，重复调用以第一次为准）
func Init(cfg Config) error {
	initOnce.Do(func() {
		defaultLogger, initErr = New(cfg)
	})
	return initErr
}

// Default 返回全局 Logger，未初始化时返回 nil
func Default() *Logger {
	return defaultLogger
}

// Sync 刷新底层缓冲，进程退出前应调用一次
func Sync() {
	if defaultLogger != nil {
		_ = defaultLogger.z.Sync()
	}
}

// Write 写入一条访问日志（ECS 规范字段名）
func (l *Logger) Write(rec Record) {
	if l == nil {
		return
	}
	fields := []zap.Field{
		zap.String("service.name", rec.Service),
		zap.String("method", rec.Method),
		zap.String("path", rec.Path),
		zap.String("query", rec.Query),
		zap.Int("status", rec.Status),
		zap.Float64("latency_ms", rec.LatencyMs),
		zap.String("client.ip", rec.ClientIP),
		zap.String("user_agent.original", rec.UserAgent),
		zap.String("request_id", rec.RequestID),
		zap.Int("bytes_out", rec.BytesOut),
	}
	if rec.Error != "" {
		fields = append(fields, zap.String("error", rec.Error))
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.z.Info("", fields...)
}
