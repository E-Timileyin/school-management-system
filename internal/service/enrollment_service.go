package service

import "github.com/E-Timileyin/school-management-system/internal/repository"

type EnrollmentService struct {
	enrollmentRepo *repository.EnrollmentRepository
}

func NewEnrollmentService(enrollmentRepo *repository.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{enrollmentRepo: enrollmentRepo}
}

func (s *EnrollmentService) DeleteEnrollment(id uint) error {
	return s.enrollmentRepo.Delete(id)
}
