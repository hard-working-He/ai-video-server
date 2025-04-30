package db

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// GetDB returns a database connection
func GetDB() *gorm.DB {
	once.Do(func() {
		dsn := "root:Henian2345..@tcp(81.68.224.194:3306)/videos_db?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		instance = db
	})
	return instance
}

// InitDB initializes the database connection and performs migrations
func InitDB() error {
	GetDB()
	log.Println("Database connection successful!")
	return nil
}
