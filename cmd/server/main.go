// cmd/server/main.go
package main

import (
	"flag"
	"log"
	"os"

	"github.com/E-Timileyin/school-management-system/internal/config"
	"github.com/E-Timileyin/school-management-system/internal/migration"
	"github.com/E-Timileyin/school-management-system/internal/routes"
)

func main() {
	// Parse command line flags
	resetDB := flag.Bool("reset-db", false, "Reset the database by dropping all tables and running migrations")
	flag.Parse()

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Reset database if flag is set
	if *resetDB {
		log.Println("Resetting database...")
		if err := migration.ResetDB(db); err != nil {
			log.Fatalf("Failed to reset database: %v", err)
		}
		log.Println("Database reset and migrations completed successfully")
	} else {
		// Run normal migrations
		if err := migration.MigrateDB(db); err != nil {
			log.Fatalf("Failed to run database migrations: %v", err)
		}
	}

	// Initialize router
	router := routes.SetupRouter(db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
