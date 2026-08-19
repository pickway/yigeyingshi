package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 连接参数 —— 从环境变量读取，未配置时给出占位符提示
// 用法示例：
//   MYSQL_HOST=127.0.0.1 MYSQL_PORT=3306 MYSQL_USER=root MYSQL_PASSWORD=secret \
//     go run scripts/init_db.go
var (
	user    = getEnv("MYSQL_USER", "MYSQL_USER_PLACEHOLDER")
	pass    = getEnv("MYSQL_PASSWORD", "MYSQL_PASSWORD_PLACEHOLDER")
	host    = getEnv("MYSQL_HOST", "MYSQL_HOST_PLACEHOLDER")
	port    = getEnv("MYSQL_PORT", "3306")
	sqlFile = "cineverse-init.sql"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

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
