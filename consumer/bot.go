package main

import (
	"encoding/csv"
	"fmt"
	"github.com/markgerald/chat-api-challenge/chatsocket"
	"github.com/markgerald/chat-api-challenge/consumer/rabbitmq"
	rediscache "github.com/markgerald/chat-api-challenge/consumer/redis"
	"log"
	"net/http"
)

func main() {
	chatHub := chatsocket.GetHub()
	ConsumeStockFromQueue(chatHub)
}

func GetStockQuote(stockCode string) (string, error) {
	resp, err := http.Get("https://stooq.com/q/l/?s=" + stockCode + "&f=sd2t2ohlcv&h&e=csv")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	price := records[1][3]
	return stockCode + " quote is $" + price + " per share", nil
}

func ConsumeStockFromQueue(hub *chatsocket.Hub) {
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			quote, err := GetStockQuote(string(d.Body))
			if err != nil {
				log.Printf("Failed to get stock quote: %s", err)

			} else {
				log.Printf("Received stock quote: %s", quote)
				rediscache.Persist(quote)
			}
		}
	}()
	fmt.Println("Waiting for messages from BOT...")
	<-forever
}
