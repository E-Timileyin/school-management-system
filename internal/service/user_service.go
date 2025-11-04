package service

import (
	"github.com/E-Timileyin/school-management-system/internal/models"
	"github.com/E-Timileyin/school-management-system/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.FindByEmail(email)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.userRepo.List()
}
