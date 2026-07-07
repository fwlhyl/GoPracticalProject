package db

import (
	"fleet-management/internal/model"
	"github.com/glebarez/sqlite" // 使用纯 Go 版本的 SQLite 驱动
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	// 连接 SQLite 数据库 (如果没有文件会自动创建 test.db)
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// 自动迁移模式 (相当于 Hibernate 的 update)
	// GORM 会自动帮我们创建表结构
	err = DB.AutoMigrate(&model.Vehicle{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database connected and migrated successfully!")
}
