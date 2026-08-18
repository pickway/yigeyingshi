package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 连接参数 —— 直接硬编码，一次性脚本
const (
	user     = "root"
	pass     = "root"
	host     = "118.145.113.88"
	port     = "3306"
	sqlFile  = "cineverse-init.sql"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&multiStatements=true&parseTime=True&loc=Local",
		user, pass, host, port)

	data, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Fatalf("读取 SQL 脚本失败: %v", err)
	}
	log.Printf("成功读取 %s (%d 字节)", sqlFile, len(data))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("打开 MySQL 连接失败: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		log.Fatalf("Ping MySQL 失败: %v", err)
	}
	log.Println("已连接到 MySQL:", host+":"+port)

	script := string(data)
	log.Println("开始执行 SQL 脚本（依赖 multiStatements）...")
	if _, err := db.Exec(script); err != nil {
		log.Fatalf("脚本执行失败: %v", err)
	}
	log.Println("全部语句执行成功 ✓")

	// 验收计数
	tables := []string{"movies", "articles", "learning_courses", "learning_paths", "ai_tools", "newsletters"}
	fmt.Println()
	fmt.Println("---------- 数据校验 ----------")
	useDB(db)
	for _, t := range tables {
		var cnt int
		if err := db.QueryRow("SELECT COUNT(*) FROM " + t).Scan(&cnt); err != nil {
			log.Fatalf("计数 %s 失败: %v", t, err)
		}
		fmt.Printf("  %-18s %d 行\n", t, cnt)
	}
	fmt.Println("-------------------------------")
}

func useDB(db *sql.DB) {
	if _, err := db.Exec("USE yigeyingshi"); err != nil {
		log.Fatalf("USE yigeyingshi 失败: %v", err)
	}
}
