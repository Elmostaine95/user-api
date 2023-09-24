package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "postgres://postgres:adminadmin@my-database.cughk4hu7go7.eu-north-1.rds.amazonaws.com:5432/user_api"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
}
