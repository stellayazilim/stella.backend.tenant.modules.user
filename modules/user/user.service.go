package user

import (
	"github.com/stellayazilim/stella.backend.tenants/entities"
	"github.com/stellayazilim/stella.backend.tenants/modules/database"
	"gorm.io/gorm"
)

type UserService interface {
	GetAll() []entities.User
	GetById(uint64) (entities.User, error)
	Create(entities.User) error
	UpdateById(uint64, entities.User) error
	DeleteById(uint64) error
}

type userService struct {
	users []entities.User
	db    *gorm.DB
	index uint64
}

func Service() UserService {
	return &userService{
		index: 0,
		db:    database.Db,
	}
}

func (s *userService) GetAll() []entities.User {

	result := []entities.User{}
	s.db.Find(&result)

	return result
}

func (s *userService) Create(u entities.User) error {

	e := s.db.Create(&u).Error
	return e
}

func (s *userService) GetById(id uint64) (entities.User, error) {

	user := entities.User{}
	err := s.db.First(&user, id).Error

	if err != nil {
		return entities.User{}, err
	}
	return user, nil

}

func (s *userService) UpdateById(id uint64, data entities.User) error {

	err := s.db.Find(&entities.User{}, id).Updates(data).Error

	return err

}

func (s *userService) DeleteById(id uint64) error {

	delUser := entities.User{}
	delUser.ID = id
	err := s.db.Delete(&delUser).Error
	return err
}
