package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type ResetPassword struct{}

func (fn ResetPassword) Consume(msgs <-chan amqp.Delivery) {
	log.Println("Into ResetPassword")
	for d := range msgs {
		log.Println("[RP - x] %s", d.Body)
	}
}

func ResetPasswordConsumer(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Println("[x] %s", d.Body)
	}
}
