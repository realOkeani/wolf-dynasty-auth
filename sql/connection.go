package sql

import (
	"github.com/realOkeani/wolf-dynasty-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:9s#2eq8bi%6r!@tcp(127.0.0.1:3306)/wolf_dynasty_db?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to DB")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
