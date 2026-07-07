package model

// Vehicle 定义了车辆的数据结构
// gorm tag 就像 Java 里的 @Column，告诉 GORM 怎么建表
type Vehicle struct {
	ID     string `json:"id" gorm:"primaryKey"` // 主键
	Brand  string `json:"brand"`
	Model  string `json:"model"`
	Status string `json:"status" gorm:"default:'idle'"` // 默认状态
}
