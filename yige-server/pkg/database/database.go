package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 使用 MySQL DSN 初始化连接
// DSN 示例: user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func Init(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	DB = db
	log.Println("[database] connected to MySQL successfully")
	return db, nil
}

func AutoMigrate(models ...interface{}) error {
	if DB == nil {
		panic("database not initialized")
	}
	log.Println("[database] running auto migration")
	return DB.AutoMigrate(models...)
}
