package entities

type Category struct {
	Base
	Name          string
	SubCategories []Category `gorm:"foreignKey:ParentId"`
	ParentId      uint64
}
