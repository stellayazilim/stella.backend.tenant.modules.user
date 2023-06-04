package repositories

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser()
}

type userRepository struct {
	db *gorm.DB
}

func CreateUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser() {

}
