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
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	port, err := strconv.Atoi(env("ADMIN_SERVER_PORT", "8081"))
	if err != nil {
		return nil, errors.New("ADMIN_SERVER_PORT 必须是数字")
	}
	cfg := &Config{
		Host: env("ADMIN_SERVER_HOST", "127.0.0.1"), Port: port,
		Mode: env("ADMIN_GIN_MODE", "debug"), DatabaseDSN: env("ADMIN_DB_DSN", "../yige-server/yige.db"),
		Username: env("ADMIN_USERNAME", "admin"), Password: env("ADMIN_PASSWORD", "change-me"),
		TokenSecret: env("ADMIN_TOKEN_SECRET", "development-secret-change-before-release"), TokenTTL: 8 * time.Hour,
		AllowedOrigin: env("ADMIN_ALLOWED_ORIGIN", "http://127.0.0.1:4174"),
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
