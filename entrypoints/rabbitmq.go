package entrypoints

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/streadway/amqp"

	"guitou.com/notification-msvc/config"
	"guitou.com/notification-msvc/models"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Request contains sending email parameter
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

// NewRequest create
func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

// SendEmail sends email
func (r *Request) SendEmail(auth smtp.Auth) (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.mailtrap.io:25"

	if err := smtp.SendMail(addr, auth, "Guitou <notification@guitou.cm>", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

// InitBroker init the connection with AMQP
func InitBroker() {
	conn, err := amqp.Dial(config.Config.RabbitMQUri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"guitou.project.user.invitated", // name
		true,                            // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			data := models.NewProjectInvitationData()
			err := data.ParseData(d.Body)
			if err != nil {
				failOnError(err, fmt.Sprintf("Unmarsheling Failed : %s", d.Body))
			}
			log.Println("DATA", data)

			auth := smtp.PlainAuth("", "39150c7f22ec69", "55963de64c6833", "smtp.mailtrap.io")
			body, err := data.ParseTemplate("templates/guitou.project.user.invited.html")
			if err == nil {
				subject := fmt.Sprintf("Join %s", data.ProjectName)
				r := NewRequest([]string{data.UserEmail}, subject, body)

				ok, _ := r.SendEmail(auth)
				log.Println(ok)
			} else {
				log.Println("ERROR occured when sending the e-mail")
				log.Println(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages from [%s]. To exit press CTRL+C", q.Name)
	<-forever
}
