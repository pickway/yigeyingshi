package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yigeyingshi/yige-server/internal/config"
	"github.com/yigeyingshi/yige-server/internal/model"
	"github.com/yigeyingshi/yige-server/internal/router"
	"github.com/yigeyingshi/yige-server/internal/service"
	"github.com/yigeyingshi/yige-server/pkg/accesslog"
	"github.com/yigeyingshi/yige-server/pkg/database"
)

func main() {
	cfg := config.Load()

	// 初始化 HTTP 访问日志（写本地 JSON 行文件，供 Filebeat 采集到 ES）
	if err := accesslog.Init(accesslog.Config{
		FilePath:   cfg.AccessLog.Path,
		MaxSizeMB:  cfg.AccessLog.MaxSizeMB,
		MaxBackups: cfg.AccessLog.MaxBackups,
		MaxAgeDays: cfg.AccessLog.MaxAgeDays,
		Compress:   cfg.AccessLog.Compress,
	}); err != nil {
		log.Fatalf("初始化 access log 失败: %v", err)
	}
	defer accesslog.Sync()

	db, err := database.Init(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 表结构已通过 scripts/cineverse-init.sql 在 MySQL 手动创建，这里仅做安全兜底（幂等）
	err = database.AutoMigrate(
		&model.Movie{},
		&model.Article{},
		&model.LearningCourse{},
		&model.LearningPath{},
		&model.AiTool{},
		&model.Newsletter{},
	)
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// 仅在 MySQL 表为空时写入种子数据（与 SQLite 版本逻辑一致）
	service.Seed(db)

	r := router.Setup(db, cfg.Server.Mode)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("[server] starting on %s", addr)

	// 优雅退出：收到信号时先 flush 访问日志再退出
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("[server] shutting down, flushing access log...")
		accesslog.Sync()
		os.Exit(0)
	}()

	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
