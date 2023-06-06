package serializers

import (
	"fmt"

	"github.com/stellayazilim/stella.backend.tenants.modules.user/entity"
)

type IUserResponse interface {
	SerializeAll(u []entity.User) *[]userResponse
	Serialize(u entity.User) *userResponse
}
type userResponse struct {
	ID        uint64   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Phone     []string `json:"phone"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
	Roles     string   `json:"roles"`
}

func CreateUserResponseSerializer() IUserResponse {
	return &userResponse{}
}

func (c *userResponse) SerializeAll(u []entity.User) *[]userResponse {
	response := []userResponse{}

	for _, r := range u {
		ur := c.Serialize(r)
		response = append(response, *ur)
	}

	return &response
}

func (r *userResponse) Serialize(u entity.User) *userResponse {
	return &userResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone, //fmt.Sprintf("%v/users/%v/contact/phone", "http://localhost:8080", r.ID),
		Email:     fmt.Sprintf("%v/users/%v/contact/email", "http://localhost:8080", r.ID),
		Roles:     fmt.Sprintf("%v/users/%v/roles", "http://localhost:8080", r.ID),
	}
}
