package services

import (
	"github.com/markgerald/chat-api-challenge/database"
	"github.com/markgerald/chat-api-challenge/dto"
	"github.com/markgerald/chat-api-challenge/models"
	"golang.org/x/crypto/bcrypt"
)

func UserDtoService(user models.User) dto.User {
	return dto.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func SaveUser(user models.User) models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	database.DB.Save(&user)
	return user
}
