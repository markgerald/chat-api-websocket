package producer

import (
	"github.com/streadway/amqp"
	"log"
)

func ConnectToRabbitMQProducer() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}
