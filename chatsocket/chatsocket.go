package chatsocket

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	redischat "github.com/markgerald/chat-api-challenge/chatsocket/redis"
	"github.com/markgerald/chat-api-challenge/consumer/db"
	"github.com/markgerald/chat-api-challenge/messages"
	"github.com/markgerald/chat-api-challenge/models"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mainHub = NewHub()

func init() {
	go mainHub.Run()
}

func GetHub() *Hub {
	return mainHub
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	send chan Message
	hub  *Hub
}

func NewClient(id string, conn *websocket.Conn, hub *Hub) *Client {
	return &Client{ID: id, Conn: conn, send: make(chan Message, 256), hub: hub}
}

func (c *Client) Read(ctx *gin.Context) {
	db.Connect()
	defer func() {
		c.hub.unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg Message
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		var messageQueue models.Message
		messageQueue.Username = msg.Sender
		messageQueue.Content = msg.Content
		ps := messages.ProcessMessage{}
		ps.ProcessMessage(ctx, messageQueue)
		c.hub.broadcast <- msg
		time.Sleep(3 * time.Second)
		if strings.HasPrefix(msg.Content, "/stock=") == true {
			var cb = context.Background()
			newContent := redischat.Client().Get(cb, "botcontent")
			msg.Sender = "bot"
			msg.Content = newContent.Val()
			log.Printf("BOT CONTENT: %s", newContent.Val())
		}
		SendMessageToClient(msg, c)
	}
}

func (c *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			} else {
				err := c.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println("Error: ", err)
					break
				}
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}

	}
}

func (c *Client) Close() {
	close(c.send)
}

type GinContext struct {
	context *gin.Context
}

func getUpgrader(ctx *gin.Context) *websocket.Conn {
	ws, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err.Error())
		return ws
	}
	return ws
}

func ServeWS(ctx *gin.Context, roomId string) {
	ws := getUpgrader(ctx)
	hub := GetHub()
	client := NewClient(roomId, ws, hub)
	hub.register <- client
	log.Println("Client registered: " + client.ID)
	go client.Write()
	go client.Read(ctx)
}

func SendMessageToClient(message Message, client *Client) {
	client.send <- message
}
