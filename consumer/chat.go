package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/markgerald/chat-api-challenge/consumer/db"
	"github.com/markgerald/chat-api-challenge/consumer/entity"
	"github.com/markgerald/chat-api-challenge/consumer/rabbitmq"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	println(os.Getenv("DBUSER"))
	db.Connect()
	ConsumeMessagesFromQueue()
}

func ConsumeMessagesFromQueue() {
	conn, err := rabbitmq.ConnectToRabbitMQ()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"messageCommands",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Waiting for messages")
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			if err != nil {
				log.Printf("Failed to get  message from chat: %s", err)
			} else {

				rawData := []byte(d.Body)
				var payload interface{}
				err := json.Unmarshal(rawData, &payload)
				if err != nil {
					log.Fatal(err)
				}
				log.Print(payload)
				m := payload.(map[string]interface{})

				var model entity.Message
				log.Print(m["sender"])
				model.Username = m["sender"].(string)
				model.Content = m["message"].(string)
				model.UserID = 1
				db.DB.Create(&model)
				log.Printf("Received message from chat: %s", string(d.Body))
			}
		}
	}()
	fmt.Println("Waiting for messages from CHAT...")
	<-forever
}
