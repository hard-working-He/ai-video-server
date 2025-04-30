package handlers

import (
	"net/http"

	"go-mysql-videos/db"
	"go-mysql-videos/models"

	"github.com/gin-gonic/gin"
)

// VideoHandler handles HTTP requests related to videos
type VideoHandler struct{}

// NewVideoHandler creates a new video handler
func NewVideoHandler() *VideoHandler {
	return &VideoHandler{}
}

// GetAllVideoLists handles the request to get all video lists
func (h *VideoHandler) GetAllVideoLists(c *gin.Context) {
	// 获取数据库连接
	dbConn := db.GetDB()

	// 声明视频列表变量
	var videoLists []models.VideoList

	// 从数据库获取所有视频列表
	result := dbConn.Find(&videoLists)
	if result.Error != nil {
		// 如果发生错误，返回500状态码
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取视频列表失败",
			"error":   result.Error.Error(),
		})
		return
	}

	// 成功获取到数据，返回结果
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "成功获取所有视频列表",
		"data":    videoLists,
		"total":   len(videoLists),
	})
}
