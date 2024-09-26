package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init(){
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Connection error:", err)
		panic(err)
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS hotel")
	if err != nil {
		log.Fatal("Error creating database:", err)
		panic(err)
	}

	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/hotel?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	Database, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		panic(err)
	}
	
}

func createSchema(db *gorm.DB) {
	
}

