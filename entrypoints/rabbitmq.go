package entrypoints

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strings"

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

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

// SendEmail sends email
func (r *Request) SendEmail(auth smtp.Auth) (bool, error) {

	// mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// msg := []byte(to + subject + mime + "\n" + r.body)
	// subject := "Subject: " + r.subject + "!\r\n"
	// to := fmt.Sprintf("To: %s\r\n", r.to[0])
	from := mail.Address{Name: "Guitou", Address: "mael.fosso@guitou.cm"}
	addr := fmt.Sprintf("%s:%s", config.GetConfig().Mail.Host, config.GetConfig().Mail.Port)

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = r.to[0]
	header["Subject"] = r.subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(r.body))

	if err := smtp.SendMail(addr, auth, from.Address, r.to, []byte(message)); err != nil {
		return false, err
	}
	return true, nil
}

// InitBroker init the connection with AMQP
func InitBroker() {
	log.Println("RabbitMQ ", config.GetConfig().RabbitMQUri)
	conn, err := amqp.Dial(config.GetConfig().RabbitMQUri)
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

			log.Println(config.GetConfig().Mail.Username, config.GetConfig().Mail.Password, config.GetConfig().Mail.Host)
			auth := smtp.PlainAuth("", config.GetConfig().Mail.Username, config.GetConfig().Mail.Password, config.GetConfig().Mail.Host)

			body, err := data.ParseTemplate("templates/guitou.project.user.invited.html")
			if err == nil {
				subject := fmt.Sprintf("Join %s", data.ProjectName)
				r := NewRequest([]string{data.UserEmail}, subject, body)

				// ok, err := r.SendEmail(auth)
				// log.Println(ok)
				// log.Println(err)

				go r.SendEmail(auth)
			} else {
				log.Println("ERROR occured when sending the e-mail")
				log.Println(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages from [%s]. To exit press CTRL+C", q.Name)
	<-forever
}
