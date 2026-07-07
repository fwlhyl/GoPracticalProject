package handler

import (
	"fleet-management/internal/db"
	"fleet-management/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVehicles 获取所有车辆列表
func GetVehicles(c *gin.Context) {
	var list []model.Vehicle
	// 相当于 SELECT * FROM vehicles;
	db.DB.Find(&list)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": list,
	})
}

// CreateVehicle 创建新车辆
func CreateVehicle(c *gin.Context) {
	var newVehicle model.Vehicle

	if err := c.ShouldBindJSON(&newVehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}

	// 相当于 INSERT INTO vehicles ...
	if err := db.DB.Create(&newVehicle).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "创建失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建成功", "data": newVehicle})
}

// GetVehicleByID 根据ID获取单辆车
func GetVehicleByID(c *gin.Context) {
	id := c.Param("id")
	var vehicle model.Vehicle

	// 相当于 SELECT * FROM vehicles WHERE id = ? LIMIT 1;
	result := db.DB.First(&vehicle, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "车辆不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": vehicle})
}

// DeleteVehicle 根据ID删除车辆
func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")

	// 相当于 DELETE FROM vehicles WHERE id = ?;
	db.DB.Delete(&model.Vehicle{}, "id = ?", id)

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})
}

// UpdateVehicleStatus 根据ID局部更新车辆状态
func UpdateVehicleStatus(c *gin.Context) {
	id := c.Param("id")
	var updateData struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "参数错误"})
		return
	}

	// 相当于 UPDATE vehicles SET status = ? WHERE id = ?;
	result := db.DB.Model(&model.Vehicle{}).Where("id = ?", id).Update("status", updateData.Status)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "车辆不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "状态更新成功"})
}
