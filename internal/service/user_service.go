package service

import (
	"whois-client/internal/models"
	"whois-client/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id uint64) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Create(user *models.User) error {
	// TODO add validation
	return s.repo.Create(user)
}

func (s *UserService) Update(id uint64, user *models.User) error {
	// TODO add validation
	return s.repo.Update(id, user)
}

func (s *UserService) Delete(id uint64) error {
	return s.repo.Delete(id)
}
