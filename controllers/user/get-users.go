package user

import (
	"github.com/gin-gonic/gin"
	"github.com/markgerald/chat-api-challenge/repository"
	"github.com/markgerald/chat-api-challenge/services"
	"net/http"
)

func GetUserById(c *gin.Context) {
	user := repository.GetUserById(c.Params.ByName("id"))
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "No User"})
		return
	}
	c.JSON(http.StatusOK, services.UserDtoService(user))
}
