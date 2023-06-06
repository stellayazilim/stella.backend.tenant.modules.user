package service

import (
	"github.com/stellayazilim/stella.backend.tenants.modules.user/database"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/entity"
	"gorm.io/gorm"
)

type UserService interface {
	GetAll() []entity.User
	GetById(uint64) (entity.User, error)
	GetByTenantIdAndEmail(data entity.User) (entity.User, bool)
	Create(entity.User) error
	UpdateById(uint64, entity.User) error
	DeleteById(uint64) error
}

type userService struct {
	users []entity.User
	db    *gorm.DB
	index uint64
}

func CreateService() UserService {
	return &userService{
		index: 0,
		db:    database.GetInstance(),
	}
}
