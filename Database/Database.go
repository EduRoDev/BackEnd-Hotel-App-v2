package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		panic(err)
	}

	dns := os.Getenv("DATABASE")
	if dns == "" {
		panic("DATABASE environment variable not set or empty")
	}

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection error:", err)
		panic(err)
	}

	fmt.Print("Connected to the database successfully\n")
	Database = db
}
