package main

import (
	"log"

	"guitou.com/notification-msvc/config"
	"guitou.com/notification-msvc/entrypoints"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	// Load configuration
	err := config.LoadConfig(".")
	failOnError(err, "An error occured when loading .env")
	log.Println(config.Config.RabbitMQUri)
	// MongoDB connection

	// RabbitMQ Connection
	entrypoints.InitBroker()

}
