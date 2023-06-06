package serializers

import (
	"github.com/stellayazilim/stella.backend.tenants.modules.user/entity"
)

type IUserCreateRequest interface {
	SerializeUserCreate() entity.User
}

func CreateUserCreateRequestSerializer() IUserCreateRequest {
	return &UserCreateRequest{}
}

type UserCreateRequest struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Phone     []string `json:"phone"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
}

func (c UserCreateRequest) SerializeUserCreate() entity.User {

	return entity.User{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Phone:     c.Phone,
		Password:  c.Password,
		TenantId:  "sdfsduhfksefj",
	}

}
