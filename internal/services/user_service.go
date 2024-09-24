package services

import "cleanarchtest/internal/models"

// UserRepository - интерфейс репозитория для взаимодействия с базой данных
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetUserById(userID string) (models.User, error)
}

// UserService - сервис для работы с пользователями
type UserService struct {
	repo UserRepository
}

// NewUserService - конструктор сервиса пользователей
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAllUsers - метод для получения всех пользователей
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

// GetAllUsers - метод для получения всех пользователей
func (s *UserService) GetUserById(userID string) (models.User, error) {
	return s.repo.GetUserById(userID)
}
