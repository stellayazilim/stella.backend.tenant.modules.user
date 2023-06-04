package entities

type Image struct {
	Base
	UserId      uint64
	Src         string
	Alt         string
	Description string
	Type        string
}
