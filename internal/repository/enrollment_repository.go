// internal/repository/enrollment_repository.go
package repository

import (
	"github.com/E-Timileyin/school-management-system/internal/models"
	"gorm.io/gorm"
)

type EnrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) *EnrollmentRepository {
	return &EnrollmentRepository{db: db}
}

func (r *EnrollmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Enrollment{}, id).Error
}
