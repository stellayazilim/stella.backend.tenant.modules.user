package entities

type User struct {
	Base
	Email      string
	Password   string
	Role       *Role
	RoleId     *uint64 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProfileImg *Image
}
