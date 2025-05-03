package routes

import (
	"go-mysql-videos/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures all the routes for the application
func SetupRouter() *gin.Engine {
	// Create a default gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源，生产环境中应限制
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Create a new video handler
	videoHandler := handlers.NewVideoHandler()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to ai-video-server!"})
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Video routes
		videos := v1.Group("/videos")
		{
			// POST接口：获取所有视频列表
			videos.POST("/videolists", videoHandler.GetAllVideoLists)
		}
	}

	return router
}
