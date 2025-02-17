
package repositories

import (
	"EcommerceAPI/internal/models"
	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	CreateUser(user models.User) error
}

// userRepository struct (implements UserRepository)
type userRepository struct {
	db *gorm.DB
}

// Constructor for userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Implement GetAllUsers
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// Implement GetUserById
func (r *userRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

// Implement CreateUser
func (r *userRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}
