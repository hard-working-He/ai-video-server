package models

import (
	"time"

	"gorm.io/gorm"
)

// JSONB is a custom type for JSON data stored in the database
type JSONB string

// VideoList defines the model structure for the video_lists table
type VideoList struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	FilePath       string    `json:"file_path" gorm:"type:varchar(255);not null"`
	CreationParams JSONB     `json:"creation_params" gorm:"type:json"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// TableName specifies the table name for the VideoList model
func (VideoList) TableName() string {
	return "video_lists"
}

// Migrate performs database migrations for the VideoList model
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&VideoList{})
}
