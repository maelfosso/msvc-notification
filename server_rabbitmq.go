package main

import "guitou.com/notification-msvc/rabbitmq"

func NewRabbitMQServer() {
	forever := make(chan bool)

	go rabbitmq.Consumer("auth.password.forget", rabbitmq.ForgetPassword{})
	go rabbitmq.Consumer("auth.password.reset", rabbitmq.ResetPassword{})
	// rabbitmq.Consumer("project.user.invited", func() {})

	<-forever
}
