package configs

import (
	"log"
	"mnewsapi/app/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error")
	} else {
		log.Println("Database connected")
	}

	/// Auto Migrate (Add field table using table berdasarkan nama type Struct)
	/// Synchronize database
	database.AutoMigrate(&models.News{})

	DB = database
	return DB
}
