package rabbitmq

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

type RabbitMQHandler interface {
	Consume(<-chan amqp.Delivery)
}

var (
	rabbitMQConn    *amqp.Connection
	rabbitMQChannel *amqp.Channel
)

func Init() {
	rabbitMQConn, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Println("Failed to connect to RabbitMQ : %s", err)
		panic(err)
	}

	rabbitMQChannel, err = rabbitMQConn.Channel()
	if err != nil {
		log.Println("Failed to open channel : %s", err)
		panic(err)
	}
}

func Close() {
	rabbitMQChannel.Close()
	rabbitMQConn.Close()
}

func Consumer(exchange string, handler RabbitMQHandler) {
	err := rabbitMQChannel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare exchange : %s", err)
		panic(err)
	}

	q, err := rabbitMQChannel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to declare a queue : %s", err)
		panic(err)
	}

	err = rabbitMQChannel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to bind a queue")
		panic(err)
	}

	messages, err := rabbitMQChannel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Failed to register a consumer")
		panic(err)
	}

	forever := make(chan bool)
	go handler.Consume(messages)
	<-forever
}
