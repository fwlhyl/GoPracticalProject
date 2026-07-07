package main

import (
	"fleet-management/internal/db"
	"fleet-management/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// 0. 初始化数据库连接
	db.InitDB()

	// 1. 初始化 Gin 引擎
	r := gin.Default()

	// 2. 定义路由组
	api := r.Group("/api/v1")
	{
		// 车辆相关的接口
		vehicles := api.Group("/vehicles")
		{
			vehicles.GET("", handler.GetVehicles)               // 获取所有车辆
			vehicles.POST("", handler.CreateVehicle)            // 创建车辆
			vehicles.GET("/:id", handler.GetVehicleByID)        // 根据ID获取单辆车
			vehicles.DELETE("/:id", handler.DeleteVehicle)      // 根据ID删除车辆
			vehicles.PATCH("/:id", handler.UpdateVehicleStatus) // 局部更新：修改车辆状态
		}
	}

	// 3. 启动服务器，监听 8081 端口
	println("🚀 Fleet Management System starting on :8081")
	err := r.Run(":8081")
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
