package service

import (
	"github.com/E-Timileyin/school-management-system/internal/models"
	"github.com/E-Timileyin/school-management-system/internal/repository"
)

type CourseService struct {
	courseRepo *repository.CourseRepository
}

func NewCourseService(courseRepo *repository.CourseRepository) *CourseService {
	return &CourseService{courseRepo: courseRepo}
}

func (s *CourseService) CreateCourse(course *models.Course) error {
	return s.courseRepo.Create(course)
}

func (s *CourseService) GetCourseByID(id uint) (*models.Course, error) {
	return s.courseRepo.FindByID(id)
}

func (s *CourseService) UpdateCourse(course *models.Course) error {
	return s.courseRepo.Update(course)
}

func (s *CourseService) DeleteCourse(id uint) error {
	return s.courseRepo.Delete(id)
}

func (s *CourseService) ListCourses() ([]models.Course, error) {
	return s.courseRepo.List()
}

func (s *CourseService) EnrollStudent(courseID, studentID uint) error {
	return s.courseRepo.EnrollStudent(courseID, studentID)
}

func (s *CourseService) RemoveEnrollment(courseID, studentID uint) error {
	return s.courseRepo.RemoveEnrollment(courseID, studentID)
}

func (s *CourseService) GetCourseStudents(courseID uint) ([]models.Enrollment, error) {
	return s.courseRepo.GetEnrollments(courseID)
}

func (s *CourseService) GetStudentEnrollments(studentID uint) ([]models.Enrollment, error) {
	return s.courseRepo.GetStudentEnrollments(studentID)
}
