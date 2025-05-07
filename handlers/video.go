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

// post一个新视频
func (h *VideoHandler) PostNewVideo(c *gin.Context) {
	// 获取数据库连接
	dbConn := db.GetDB()

	// 定义请求参数结构
	type VideoRequest struct {
		TaskID         string `json:"task_id" binding:"required"`
		CreationParams string `json:"creation_params" binding:"required"`
		Status         string `json:"status" binding:"required"`
	}

	// 解析请求参数
	var req VideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 检查taskid是否已存在
	var existingVideo models.VideoList
	result := dbConn.Where("task_id = ?", req.TaskID).First(&existingVideo)
	if result.Error == nil {
		// 如果没有错误，说明找到了记录，返回已存在的提示
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "视频任务已存在",
			"data":    existingVideo,
		})
		return
	}

	// 创建视频对象
	video := models.VideoList{
		TaskID:         req.TaskID,
		CreationParams: req.CreationParams,
		Status:         req.Status,
		FilePath:       "", // 初始为空，后续会通过UpdateAIVideo更新
	}

	// 将视频数据保存到数据库
	result = dbConn.Create(&video)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "保存视频失败",
			"error":   result.Error.Error(),
		})
		return
	}

	// 保存成功，返回200状态码
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "视频保存成功",
		"data":    video,
	})
}

// UpdateAIVideo handles the request to update a video's status and file path
func (h *VideoHandler) UpdateAIVideo(c *gin.Context) {
	// 获取数据库连接
	dbConn := db.GetDB()

	// 定义请求参数结构
	type UpdateRequest struct {
		TaskID   string `json:"task_id" binding:"required"`
		Status   string `json:"status" binding:"required"`
		FilePath string `json:"file_path" binding:"required"`
	}

	// 解析请求参数
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 根据TaskID查找视频记录
	var video models.VideoList
	result := dbConn.Where("task_id = ?", req.TaskID).First(&video)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "未找到对应的视频记录",
			"error":   result.Error.Error(),
		})
		return
	}

	// 更新视频状态和文件路径
	video.Status = req.Status
	video.FilePath = req.FilePath

	// 保存更新
	result = dbConn.Save(&video)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "更新视频失败",
			"error":   result.Error.Error(),
		})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "视频更新成功",
		"data":    video,
	})
}
