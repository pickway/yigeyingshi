package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
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

var AppConfig *Config

func Load() *Config {
	_ = godotenv.Load()

	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))

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
	}

	log.Printf("[config] loaded: server=%s:%d db=%s", AppConfig.Server.Host, AppConfig.Server.Port, AppConfig.Database.Driver)
	return AppConfig
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
