package producer

import (
	"github.com/streadway/amqp"
	"log"
	"os"
)

func ConnectToRabbitMQProducer() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}
