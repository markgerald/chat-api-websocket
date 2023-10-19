package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Successfully connected to RabbitMQ")
	return conn, nil
}
