package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/markgerald/chat-api-challenge/chatsocket"
	"github.com/markgerald/chat-api-challenge/controllers/auth"
	"github.com/markgerald/chat-api-challenge/controllers/chat"
	"github.com/markgerald/chat-api-challenge/controllers/user"
	"github.com/markgerald/chat-api-challenge/middlewares"
)

func HandleRequests() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))
	r.POST("/register", user.Register)
	r.POST("/login", auth.Login)
	r.GET("/ws", func(c *gin.Context) {
		chatsocket.ServeWS(c, "1")
	})
	protected := r.Group("/")

	protected.Use(middlewares.JwtMiddleware())
	{
		protected.GET("/messages", chat.Chat{}.GetLastMessages)
		//protected.POST("/messages", chat.PostMessage)
	}

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
