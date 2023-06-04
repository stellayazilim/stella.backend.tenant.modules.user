package entities

type User struct {
	Base
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
