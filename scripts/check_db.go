package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 连接参数 —— 从环境变量读取，用法同 init_db.go
var (
	user = getEnv("MYSQL_USER", "MYSQL_USER_PLACEHOLDER")
	pass = getEnv("MYSQL_PASSWORD", "MYSQL_PASSWORD_PLACEHOLDER")
	host = getEnv("MYSQL_HOST", "MYSQL_HOST_PLACEHOLDER")
	port = getEnv("MYSQL_PORT", "3306")
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("打开 MySQL 连接失败: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Ping MySQL 失败: %v", err)
	}
	log.Println("已连接到 MySQL:", host+":"+port)

	fmt.Println("\n=== SHOW DATABASES ===")
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatalf("SHOW DATABASES 失败: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalf("Scan 失败: %v", err)
		}
		fmt.Printf("  - %s\n", name)
	}

	fmt.Println("\n=== 当前连接身份 ===")
	var userVal, hostVal string
	_ = db.QueryRow("SELECT CURRENT_USER(), @@hostname").Scan(&userVal, &hostVal)
	fmt.Printf("  CURRENT_USER() = %s\n  hostname = %s\n", userVal, hostVal)

	fmt.Println("\n=== SHOW GRANTS ===")
	grRows, err := db.Query("SHOW GRANTS")
	if err != nil {
		fmt.Printf("  SHOW GRANTS 失败: %v\n", err)
	} else {
		for grRows.Next() {
			var g string
			_ = grRows.Scan(&g)
			fmt.Printf("  - %s\n", g)
		}
		grRows.Close()
	}

	fmt.Println("\n=== 尝试 CREATE DATABASE 测试 ===")
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS yigeyingshi_probe CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci")
	if err != nil {
		fmt.Printf("  CREATE 失败: %v\n", err)
	} else {
		fmt.Println("  CREATE 成功 ✓ —— 权限其实没问题")
		_, _ = db.Exec("DROP DATABASE IF EXISTS yigeyingshi_probe")
	}
}

