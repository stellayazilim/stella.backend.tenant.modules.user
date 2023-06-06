package service

import "github.com/stellayazilim/stella.backend.tenants.modules.user/entity"

func (s *userService) UpdateById(id uint64, data entity.User) error {

	err := s.db.Find(&entity.User{}, id).Updates(data).Error

	return err

}
