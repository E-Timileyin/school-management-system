package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the global database connection
var DB *gorm.DB

// ConnectDB initializes and connects to the PostgreSQL database
func ConnectDB() error {
	// Load environment variables
	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("Could not load .env.local, falling back to system environment")
	}

	// Read DB variables from env
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	// Construct DSN (connection string)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSL,
	)

	// Try connecting
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	log.Println("Connected to database successfully")
	return nil
}
