package user

type userResponse struct {
	ID uint `json:"-"`
}

func (self *UserResponse) Response() UserResponse {

	return &UserResponse{}
}
