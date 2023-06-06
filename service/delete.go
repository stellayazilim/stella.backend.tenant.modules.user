package service

import "github.com/stellayazilim/stella.backend.tenants.modules.user/entity"

func (s *userService) DeleteById(id uint64) error {

	delUser := entity.User{}
	delUser.ID = id
	err := s.db.Delete(&delUser).Error
	return err
}
