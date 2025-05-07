package models

import (
	"time"

	"gorm.io/gorm"
)

// VideoList 定义了 video_lists 表的模型结构
type VideoList struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// 主键 ID，在 JSON 中显示为 "id"，在数据库中作为主键
	FilePath       string    `json:"file_path" gorm:"column:file_path;type:varchar(255);not null"`    // 文件路径，不能为空，在数据库中设置为 varchar(255) 类型
	CreationImg    string    `json:"creation_img" gorm:"column:creation_img;type:varchar(255)"`       // 创建图片，普通字符串类型
	CreationParams string    `json:"creation_params" gorm:"column:creation_params;type:varchar(255)"` // 创建参数，普通字符串类型
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`                                // 创建时间，使用 GORM 自动填充当前时间
	Status         string    `json:"status" gorm:"column:status;type:varchar(255);not null"`
	TaskID         string    `json:"task_id" gorm:"column:task_id;type:varchar(255);not null"` // 新增字段：任务ID，不能为空
}

// TableName 指定 VideoList 模型对应的数据库表名
// 这个方法会被 GORM 调用以确定表名
func (VideoList) TableName() string {
	return "video_lists" // 返回表名为 "video_lists"
}

// Migrate 为 VideoList 模型执行数据库迁移
// 此函数接收一个 GORM 数据库连接，并使用它创建或更新数据库表结构
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&VideoList{}) // 自动迁移 VideoList 模型到数据库，返回可能出现的错误
}
