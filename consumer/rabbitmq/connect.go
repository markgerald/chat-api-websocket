package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Successfully connected to RabbitMQ")
	return conn, nil
}
