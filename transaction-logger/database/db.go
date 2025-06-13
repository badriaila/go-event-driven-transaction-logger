package database

import (
	"fmt"
	"log"
	"os"

	"transaction-logger/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB initializes the database connection

var DB *gorm.DB

func ConnectDB() {
	_ = godotenv.Load(".env")

	dsn := fmt.Sprintf( // Data Source Name (DSN) for PostgreSQL
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Open a connection to the PostgreSQL database

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Transaction{}) // Automatically migrate the Transaction model to the database schema
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	DB = db
	log.Println("Database connection established successfully")

	return
}
