package user

import (
	"fmt"

	"github.com/stellayazilim/stella.backend.tenants/entities"
)

type UserService interface {
	GetAll() []entities.User
	GetById(uint64) (entities.User, error)
	Save(entities.User)
	UpdateById(uint64, entities.User) error
	DeleteById(uint64) error
}

type userService struct {
	users []entities.User
	index uint64
}

func Service() UserService {
	return &userService{
		index: 0,
	}
}

func (s *userService) GetAll() []entities.User {
	return s.users
}

func (s *userService) Save(u entities.User) {

	u.Id = s.index
	s.index++
	s.users = append(s.users, u)
}

func (s *userService) GetById(id uint64) (entities.User, error) {

	for _, v := range s.users {

		if v.Id == id {
			return v, nil
		}
	}

	return entities.User{}, fmt.Errorf("User not found")

}

func (s *userService) UpdateById(id uint64, data entities.User) error {

	for _, v := range s.users {

		if v.Id == id {
			s.users[id].Name = data.Name
			return nil
		}
	}

	return fmt.Errorf("not found")

}

func (s *userService) DeleteById(id uint64) error {
	for _, v := range s.users {

		if v.Id == id {
			s.users = append(s.users[:id], s.users[id+1:]...)
			return nil
		}
	}
	return fmt.Errorf("not found")
}
