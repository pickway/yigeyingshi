package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	DB = db
	log.Println("[database] connected successfully")
	return db, nil
}

func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		panic("database not initialized")
	}
	log.Println("[database] running auto migration")
	return DB.AutoMigrate(models...)
}
