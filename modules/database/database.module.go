package database

import (
	"github.com/stellayazilim/stella.backend.tenants/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// will be moved here, dw about exposed credentials
var dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Istanbul"

var db *gorm.DB

func InitalizeDatabase() {
	_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	_db.AutoMigrate(&entities.User{})
	_db.AutoMigrate(&entities.Category{})
	_db.AutoMigrate(&entities.Image{})
	_db.AutoMigrate(&entities.Role{})

	if err != nil {
		panic(err)
	}
	db = _db
}

func GetInstance() *gorm.DB {
	return db
}
