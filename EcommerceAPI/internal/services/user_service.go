package services

import (
	"EcommerceAPI/internal/models"
	"EcommerceAPI/internal/repositories"
)

type UserService interface{
	GetAllUsers() ([]models.User, error)
	GetUserById(int uint) (models.User, error)
	CreateUser(user models.User) error
}

type userService struct{
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService{
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error){
	return s.repo.GetAllUsers()
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserById(id uint) (models.User, error) {
	return s.repo.GetUserById(id)
}

// CreateUser handles user creation
func (s *userService) CreateUser(user models.User) error {
	// Business logic (e.g., email validation) can be added here
	return s.repo.CreateUser(user)
}
