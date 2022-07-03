package database

import (
	"fibr/controllers/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	connection, err := gorm.Open(mysql.Open(dsn: "sqluser:password/go_auth"), &gorm.Config{})

	if err != nil {
		panic(v: "could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
