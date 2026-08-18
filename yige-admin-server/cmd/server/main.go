package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yigeyingshi/yige-admin-server/internal/config"
	"github.com/yigeyingshi/yige-admin-server/internal/database"
	"github.com/yigeyingshi/yige-admin-server/internal/router"
	"github.com/yigeyingshi/yige-admin-server/pkg/accesslog"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

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

	db, err := database.Open(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	engine := router.Setup(db, cfg, cfg.Mode)
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("管理服务启动于 %s", address)

	// 优雅退出：收到信号时先 flush 访问日志再退出
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("[admin] shutting down, flushing access log...")
		accesslog.Sync()
		os.Exit(0)
	}()

	if err := engine.Run(address); err != nil {
		log.Fatal(err)
	}
}
