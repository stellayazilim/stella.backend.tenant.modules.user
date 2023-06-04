package dto

type UserCreateDto struct {
	Email    string `json:"email" binding:"required email max=150"`
	Password string `json:"password" binding:"required min=8 max=150"`
}
