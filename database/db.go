package database

import (
	"github.com/markgerald/chat-api-challenge/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	dsn := os.Getenv("DBUSER") + ":" + os.Getenv("DBPASSWORD") +
		"@tcp(" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + ")/" + os.Getenv("DBNAME") + "?utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Panic("Error connecting database " + err.Error())
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
}
