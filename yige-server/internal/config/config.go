package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	AccessLog AccessLogConfig
}

type ServerConfig struct {
	Host string
	Port int
	Mode string
}

type DatabaseConfig struct {
	Driver string
	DSN    string
}

// AccessLogConfig HTTP 访问日志（写本地文件，供 Filebeat 采集）
type AccessLogConfig struct {
	Path       string // 日志文件路径
	MaxSizeMB  int    // 单文件最大 MB
	MaxBackups int    // 保留旧文件数
	MaxAgeDays int    // 旧文件保留天数
	Compress   bool   // 是否 gzip 压缩旧文件
}

var AppConfig *Config

func Load() *Config {
	_ = godotenv.Load()

	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	maxSize, _ := strconv.Atoi(getEnv("ACCESS_LOG_MAX_SIZE_MB", "100"))
	maxBackups, _ := strconv.Atoi(getEnv("ACCESS_LOG_MAX_BACKUPS", "5"))
	maxAge, _ := strconv.Atoi(getEnv("ACCESS_LOG_MAX_AGE_DAYS", "7"))

	AppConfig = &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: port,
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Driver: getEnv("DB_DRIVER", "sqlite"),
			DSN:    getEnv("DB_DSN", "./yige.db"),
		},
		AccessLog: AccessLogConfig{
			Path:       getEnv("ACCESS_LOG_PATH", "logs/access.log"),
			MaxSizeMB:  maxSize,
			MaxBackups: maxBackups,
			MaxAgeDays: maxAge,
			Compress:   getEnv("ACCESS_LOG_COMPRESS", "true") == "true",
		},
	}

	log.Printf("[config] loaded: server=%s:%d db=%s access_log=%s", AppConfig.Server.Host, AppConfig.Server.Port, AppConfig.Database.Driver, AppConfig.AccessLog.Path)
	return AppConfig
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
