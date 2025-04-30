package main

import (
	"log"

	"go-mysql-videos/db"
	"go-mysql-videos/models"
	"go-mysql-videos/routes"
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

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
