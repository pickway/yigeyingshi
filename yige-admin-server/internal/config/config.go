package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Host          string
	Port          int
	Mode          string
	DatabaseDSN   string
	Username      string
	Password      string
	TokenSecret   string
	TokenTTL      time.Duration
	AllowedOrigin string
	AccessLog     AccessLogConfig
}

// AccessLogConfig HTTP 访问日志（写本地文件，供 Filebeat 采集）
type AccessLogConfig struct {
	Path       string // 日志文件路径
	MaxSizeMB  int    // 单文件最大 MB
	MaxBackups int    // 保留旧文件数
	MaxAgeDays int    // 旧文件保留天数
	Compress   bool   // 是否 gzip 压缩旧文件
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	port, err := strconv.Atoi(env("ADMIN_SERVER_PORT", "8081"))
	if err != nil {
		return nil, errors.New("ADMIN_SERVER_PORT 必须是数字")
	}
	maxSize, _ := strconv.Atoi(env("ADMIN_ACCESS_LOG_MAX_SIZE_MB", "100"))
	maxBackups, _ := strconv.Atoi(env("ADMIN_ACCESS_LOG_MAX_BACKUPS", "5"))
	maxAge, _ := strconv.Atoi(env("ADMIN_ACCESS_LOG_MAX_AGE_DAYS", "7"))
	cfg := &Config{
		Host: env("ADMIN_SERVER_HOST", "127.0.0.1"), Port: port,
		Mode: env("ADMIN_GIN_MODE", "debug"), DatabaseDSN: env("ADMIN_DB_DSN", "../yige-server/yige.db"),
		Username: env("ADMIN_USERNAME", "admin"), Password: env("ADMIN_PASSWORD", "change-me"),
		TokenSecret: env("ADMIN_TOKEN_SECRET", "development-secret-change-before-release"), TokenTTL: 8 * time.Hour,
		AllowedOrigin: env("ADMIN_ALLOWED_ORIGIN", "http://127.0.0.1:4174"),
		AccessLog: AccessLogConfig{
			Path:       env("ADMIN_ACCESS_LOG_PATH", "logs/access.log"),
			MaxSizeMB:  maxSize,
			MaxBackups: maxBackups,
			MaxAgeDays: maxAge,
			Compress:   env("ADMIN_ACCESS_LOG_COMPRESS", "true") == "true",
		},
	}
	if cfg.Mode == "release" && (cfg.Password == "change-me" || len(cfg.TokenSecret) < 32) {
		return nil, errors.New("release 模式必须配置安全的管理员密码和至少 32 位令牌密钥")
	}
	return cfg, nil
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
