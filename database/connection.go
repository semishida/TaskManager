package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"task/models"
)

var Db *gorm.DB

func InitDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		panic(err)
	}
	dsn := os.Getenv("DB_CONFIG")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database")
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection success!")

	Db.AutoMigrate(&models.Task{}, &models.User{})

}
