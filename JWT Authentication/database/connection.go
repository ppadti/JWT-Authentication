package database

import (
	"github.com/pushpa/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:password/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	 
	if err != nil {
		panic("could not connect to database")
	}
	DB = connection

	connection.AutoMigrate(&models.User{})
}