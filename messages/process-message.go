package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/chat-api-challenge/models"
	"github.com/markgerald/chat-api-challenge/producer/producermessage"
	"github.com/markgerald/chat-api-challenge/producer/producerstock"
	"log"
	"strings"
)

type ProcessMessage struct{}

func (pm *ProcessMessage) ProcessMessage(c *gin.Context, message models.Message) models.Message {
	if strings.HasPrefix(message.Content, "/stock=") == true {
		log.Printf("OI ENTREI AQUI!")
		stockCode := strings.TrimPrefix(message.Content, "/stock=")
		producerstock.SendStockToQueue([]byte(stockCode))
		message.Content = "Fetching stock data..."
		message.Username = "bot"
		return message

	}
	newMessage := `{"sender": "` + message.Username + `", "message":"` + message.Content + `"}`
	producermessage.SendToQueueMessage([]byte(newMessage), c)
	log.Printf("AH mensagem tá seno mandada assim: " + newMessage)
	c.Set("message", message)
	c.Next()
	return message

}
