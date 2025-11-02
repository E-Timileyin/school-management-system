package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the global database connection
var DB *gorm.DB

// ConnectDB initializes and connects to PostgreSQL
func ConnectDB() error {
	// Get DATABASE_URL from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	DB = db
	log.Println("✅ Connected to database successfully")
	return nil
}
