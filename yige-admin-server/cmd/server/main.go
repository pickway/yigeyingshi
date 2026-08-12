package main

import (
	"fmt"
	"log"

	"github.com/yigeyingshi/yige-admin-server/internal/config"
	"github.com/yigeyingshi/yige-admin-server/internal/database"
	"github.com/yigeyingshi/yige-admin-server/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Open(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	engine := router.Setup(db, cfg, cfg.Mode)
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("管理服务启动于 %s", address)
	if err := engine.Run(address); err != nil {
		log.Fatal(err)
	}
}
