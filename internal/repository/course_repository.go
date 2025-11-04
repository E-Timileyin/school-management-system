package repository

import (
	"github.com/E-Timileyin/school-management-system/internal/models"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *CourseRepository) FindByID(id uint) (*models.Course, error) {
	var course models.Course
	err := r.db.Preload("Teacher").First(&course, id).Error
	return &course, err
}

func (r *CourseRepository) Update(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *CourseRepository) Delete(id uint) error {
	return r.db.Delete(&models.Course{}, id).Error
}

func (r *CourseRepository) List() ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Preload("Teacher").Find(&courses).Error
	return courses, err
}

func (r *CourseRepository) EnrollStudent(courseID, studentID uint) error {
	enrollment := models.Enrollment{
		StudentID: studentID,
		CourseID:  courseID,
	}
	return r.db.Create(&enrollment).Error
}

func (r *CourseRepository) RemoveEnrollment(courseID, studentID uint) error {
	return r.db.Where("course_id = ? AND student_id = ?", courseID, studentID).
		Delete(&models.Enrollment{}).Error
}

func (r *CourseRepository) GetEnrollments(courseID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Preload("Student").
		Where("course_id = ?", courseID).
		Find(&enrollments).Error
	return enrollments, err
}

func (r *CourseRepository) GetStudentEnrollments(studentID uint) ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	err := r.db.Preload("Course").
		Where("student_id = ?", studentID).
		Find(&enrollments).Error
	return enrollments, err
}
