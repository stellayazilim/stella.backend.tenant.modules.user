package service

import "github.com/stellayazilim/stella.backend.tenants.modules.user/entity"

func (s *userService) GetAll() []entity.User {

	result := []entity.User{}
	s.db.Find(&result)

	return result
}
