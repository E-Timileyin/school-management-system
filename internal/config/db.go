package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Loader ( db connection )
func LoadDBConfig() (DatabaseConfig, error) {
	if err := godotenv.Load(".env.local"); err != nil {
		return DatabaseConfig{}, fmt.Errorf("failed to load .env file: %w", err)
	}

	if err := godotenv.Overload("../../.env.local"); err == nil {
		log.Println("Loaded environment variables from .env.local")
	}

	config := DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	if config.Host == "" || config.Port == "" || config.User == "" || config.DBName == "" {
		return DatabaseConfig{}, fmt.Errorf("missing required database configuration")
	}

	if config.SSLMode == "" {
		config.SSLMode = "disable"
	}

	return config, nil
}

func InitDB() (*gorm.DB, error) {
	config, err := LoadDBConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
	})

	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
	log.Println("Database connection established")
	return db, nil
}
