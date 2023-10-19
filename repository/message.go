package repository

import (
	"github.com/markgerald/chat-api-challenge/database"
	"github.com/markgerald/chat-api-challenge/models"
)

func GetLastMessages(quantity int) []models.Message {
	var messages []models.Message
	database.DB.Order("created_at desc").Limit(quantity).Find(&messages)
	return messages
}
