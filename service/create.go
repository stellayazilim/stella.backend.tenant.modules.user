package service

import "github.com/stellayazilim/stella.backend.tenants.modules.user/entity"

func (s *userService) Create(u entity.User) error {

	e := s.db.Create(&u).Error
	return e
}
