# Chat Websocket

## Start
Before running the go binaries, run docker that contains the RabbitMQ database and queues:

    docker-compose up -d


This project has one main application. and 2 consumers/brokers
To run the main application:

    go run main. Go

To run consumers (on separate terminals):

    go run the consumer/bot

    go run consumer/chat