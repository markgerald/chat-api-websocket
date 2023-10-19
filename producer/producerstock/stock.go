package producerstock

import (
	"github.com/markgerald/chat-api-challenge/producer"
	"github.com/streadway/amqp"
	"log"
)

func SendStockToQueue(data []byte) {
	conn, err := producer.ConnectToRabbitMQProducer()
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
		"stockCommands",
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

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}
