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

func init() {
	// Load configuration
	err := config.LoadConfig(".env")
	failOnError(err, "An error occured when loading .env")
}

func main() {

	// RabbitMQ Connection
	entrypoints.InitBroker()

}
