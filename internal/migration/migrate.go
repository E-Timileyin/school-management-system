// Package migration provides database schema migration functionality
// It handles the creation, updates, and resets of database tables and their relationships
package migration

import (
	"fmt"
	"log"

	"github.com/E-Timileyin/school-management-system/internal/models"
	"gorm.io/gorm"
)

// MigrateDB performs all necessary database migrations
// It creates tables, adds foreign keys, and sets up necessary database extensions
// db: A pointer to the gorm.DB instance to run migrations against
// Returns an error if any migration step fails
func MigrateDB(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Enable UUID extension for PostgreSQL if it doesn't exist
	// This is required for generating UUID primary keys
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Printf("Warning: Could not create uuid-ossp extension: %v", err)
		// Continue with migrations even if extension creation fails
		// as it might already exist or not be needed
	}

	// AutoMigrate creates tables and adds missing columns, but won't change column types
	// or delete unused columns to protect your data
	err := db.AutoMigrate(
		// Core authentication and user management
		&models.User{},      // Base user model with authentication details
		&models.Student{},   // Student-specific information (extends User)
		&models.Teacher{},   // Teacher-specific information (extends User)
		&models.Course{},    // Course information
		&models.Enrollment{}, // Student-course enrollment records
	)

	if err != nil {
		return fmt.Errorf("failed to auto-migrate database: %v", err)
	}

	// SQL to add foreign key constraints if they don't already exist
	// Using PL/pgSQL anonymous code block to conditionally add constraints
	sql := `
-- Begin transaction block
DO $$
BEGIN
	-- Ensure students.user_id references users.id with CASCADE delete
	-- This ensures referential integrity between students and users
	IF NOT EXISTS (
		SELECT 1
		FROM   information_schema.table_constraints
		WHERE  constraint_type = 'FOREIGN KEY'
		AND    table_name = 'students'
		AND    constraint_name = 'fk_students_user'
	) THEN
		ALTER TABLE students
		ADD CONSTRAINT fk_students_user
		FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE;  -- Delete student when user is deleted
	END IF;

	-- Ensure teachers.user_id references users.id with CASCADE delete
	-- Maintains referential integrity between teachers and users
	IF NOT EXISTS (
		SELECT 1
		FROM   information_schema.table_constraints
		WHERE  constraint_type = 'FOREIGN KEY'
		AND    table_name = 'teachers'
		AND    constraint_name = 'fk_teachers_user'
	) THEN
		ALTER TABLE teachers
		ADD CONSTRAINT fk_teachers_user
		FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE;  -- Delete teacher when user is deleted
	END IF;

	-- Ensure enrollments.student_id references students.id
	-- Maintains relationship between enrollments and students
	IF NOT EXISTS (
		SELECT 1
		FROM   information_schema.table_constraints
		WHERE  constraint_type = 'FOREIGN KEY'
		AND    table_name = 'enrollments'
		AND    constraint_name = 'fk_enrollments_student'
	) THEN
		ALTER TABLE enrollments
		ADD CONSTRAINT fk_enrollments_student
		FOREIGN KEY (student_id) REFERENCES students(id)
		ON DELETE CASCADE;  -- Delete enrollment if student is deleted
	END IF;

	-- Ensure enrollments.course_id references courses.id
	-- Maintains relationship between enrollments and courses
	IF NOT EXISTS (
		SELECT 1
		FROM   information_schema.table_constraints
		WHERE  constraint_type = 'FOREIGN KEY'
		AND    table_name = 'enrollments'
		AND    constraint_name = 'fk_enrollments_course'
	) THEN
		ALTER TABLE enrollments
		ADD CONSTRAINT fk_enrollments_course
		FOREIGN KEY (course_id) REFERENCES courses(id)
		ON DELETE CASCADE;  -- Delete enrollment if course is deleted
	END IF;

	-- Ensure courses.teacher_id references teachers.id
	-- Links courses to their respective teachers
	IF NOT EXISTS (
		SELECT 1
		FROM   information_schema.table_constraints
		WHERE  constraint_type = 'FOREIGN KEY'
		AND    table_name = 'courses'
		AND    constraint_name = 'fk_courses_teacher'
	) THEN
		ALTER TABLE courses
		ADD CONSTRAINT fk_courses_teacher
		FOREIGN KEY (teacher_id) REFERENCES teachers(id)
		ON DELETE SET NULL;  /* Set teacher_id to NULL if teacher is deleted */
	END IF;
END $$;`  /* End of transaction block */

	// Execute the SQL to add constraints
	// Note: We log but don't fail the entire migration if constraints can't be added
	// as they might already exist or the database user might not have sufficient permissions
	if err := db.Exec(sql).Error; err != nil {
		log.Printf("Warning: Could not add foreign key constraints: %v", err)
		// Continue with the migration even if constraints can't be added
	}

	log.Println("Database migrations completed successfully")
	return nil
}
