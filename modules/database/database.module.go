package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// will be moved here, dw about exposed credentials
var dsn = "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Istanbul"

var Db *gorm.DB

func DatabaseModule() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	Db = db
}
