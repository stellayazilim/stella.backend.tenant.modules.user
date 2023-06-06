package service

import "github.com/stellayazilim/stella.backend.tenants.modules.user/entity"

func (s *userService) GetById(id uint64) (entity.User, error) {

	user := entity.User{}
	err := s.db.First(&user, id).Error

	if err != nil {
		return entity.User{}, err
	}
	return user, nil

}

func (s *userService) GetByTenantIdAndEmail(data entity.User) (entity.User, bool) {
	user := entity.User{}
	q := s.db.Find(&user, &entity.User{TenantId: data.TenantId, Email: data.Email}).Limit(1)

	if q.RowsAffected > 0 {
		return user, true
	}
	return user, false

}
