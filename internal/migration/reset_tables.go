package migration

import (
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

// ResetDB drops all tables and runs migrations from scratch
// WARNING: This will delete all data in the database
func ResetDB(db *gorm.DB) error {
	// Get all tables
	var tables []string
	if err := db.Raw(`
		SELECT tablename 
		FROM pg_tables 
		WHERE schemaname = 'public'
		AND tablename != 'schema_migrations'
	`).Scan(&tables).Error; err != nil {
		return fmt.Errorf("failed to get list of tables: %w", err)
	}

	if len(tables) == 0 {
		log.Println("No tables to drop")
	} else {
		// Drop all tables with CASCADE to handle foreign key constraints
		sql := fmt.Sprintf(`DROP TABLE IF EXISTS %s CASCADE`, strings.Join(tables, ", "))
		if err := db.Exec(sql).Error; err != nil {
			return fmt.Errorf("failed to drop tables: %w", err)
		}
		log.Printf("Dropped %d tables\n", len(tables))
	}

	// Run migrations to recreate the schema
	log.Println("Running migrations...")
	if err := MigrateDB(db); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database reset and migrations completed successfully")
	return nil
}
