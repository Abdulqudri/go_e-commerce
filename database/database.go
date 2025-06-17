package database

import (
	"log"

	"github.com/Abdulqudri/myapi/configs"
	"github.com/Abdulqudri/myapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase() {
	dsn := configs.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	DB = db

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration error:", err)
	}
	log.Println("Connected to MySQL & migrated User model")
}
