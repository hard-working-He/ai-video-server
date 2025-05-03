package main

import (
	"log"

	"go-mysql-videos/db"
	"go-mysql-videos/models"
	"go-mysql-videos/routes"

	"github.com/gin-contrib/cors"
)

func main() {
	// Initialize the database connection
	if err := db.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Perform database migrations
	if err := models.Migrate(db.GetDB()); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Setup the router
	r := routes.SetupRouter()

	// 使用默认CORS配置
	r.Use(cors.Default())

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
