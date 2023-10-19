package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/chat-api-challenge/repository"
)

type Chat struct {
}

func (chat Chat) GetLastMessages(c *gin.Context) {
	messages := repository.GetLastMessages(50)
	c.JSON(200, messages)
}
