package db

import (
	"fmt"
	"os"

	"fleet-management/internal/model"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
// 通过环境变量 DB_DRIVER 切换数据库：
//
//	不设置或 "sqlite" → 使用 SQLite（test.db）
//	"mysql"           → 使用 MySQL（需同时设置 DB_DSN）
func InitDB() {
	var err error
	driver := os.Getenv("DB_DRIVER")

	if driver == "mysql" {
		dsn := os.Getenv("DB_DSN")
		if dsn == "" {
			log.Fatal("DB_DRIVER=mysql 时必须设置 DB_DSN，例如: root:password@tcp(127.0.0.1:3306)/fleet?charset=utf8mb4&parseTime=True")
		}
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		// 默认使用 SQLite
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// 自动迁移（相当于 Hibernate 的 ddl-auto: update）
	err = DB.AutoMigrate(&model.Vehicle{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	fmt.Printf("Database connected! driver=%s\n", driver)
}
