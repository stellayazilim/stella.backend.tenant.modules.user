package entity

import (
	"time"

	pq "github.com/lib/pq"
	"gorm.io/gorm"
)

type PhoneNumbers []string
type User struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TenantId  string
	FirstName string         `gorm:"type:varchar(25)"`
	LastName  string         `gorm:"type:varchar(25)"`
	Email     string         `gorm:"type:varchar(50)"`
	Phone     pq.StringArray `gorm:"type:varchar(24)[]"`
	Password  string         `gorm:"type:varchar(50)"`
	Roles     []byte
}
