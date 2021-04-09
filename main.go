package main

import (
	"guitou.com/notification-msvc/rabbitmq"
)

// func failOnError(err error, msg string) {
// 	if err != nil {
// 		log.Fatalf("%s: %s", msg, err)
// 	}
// }

// func init() {
// 	// Load configuration
// 	err := config.LoadConfig(".env")
// 	failOnError(err, "An error occured when loading .env")
// }

func main() {

	forever := make(chan bool)
	go func() {
		rabbitmq.Init()
		NewRabbitMQServer()

		defer rabbitmq.Close()
	}()

	go func() {
		// For http Request
	}()

	<-forever
}
