package user

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/chat-api-challenge/models"
	"github.com/markgerald/chat-api-challenge/services"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	savedUser := services.SaveUser(user)
	c.JSON(200, gin.H{"message": "Successfully registered", "user": savedUser})
}
