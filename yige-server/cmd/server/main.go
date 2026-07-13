package main

import (
	"fmt"
	"log"

	"github.com/yigeyingshi/yige-server/internal/config"
	"github.com/yigeyingshi/yige-server/internal/model"
	"github.com/yigeyingshi/yige-server/internal/router"
	"github.com/yigeyingshi/yige-server/internal/service"
	"github.com/yigeyingshi/yige-server/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.Init(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

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

	service.Seed(db)

	r := router.Setup(db, cfg.Server.Mode)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("[server] starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
