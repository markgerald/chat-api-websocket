package repository

import (
	"github.com/markgerald/chat-api-challenge/database"
	"github.com/markgerald/chat-api-challenge/models"
)

func GetUserById(id string) models.User {
	var user models.User
	database.DB.Preload("User").First(&user, id)
	return user
}

func GetUserByEmail(email string, user models.User) models.User {
	database.DB.Where("email = ?", email).First(&user)
	return user
}
