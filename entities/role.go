package entities

type Role struct {
	Base
	Users []User
	Name  string
	Perms []uint8
}
