package rabbitmq

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ForgetPassword struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Token     string             `json:"token" bson:"token"`
	ExpiredAt time.Time          `json:"expiredAt,omitempty" bson:"expiredAt"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

func (fp ForgetPassword) Consume(msgs <-chan amqp.Delivery) {
	log.Println("Into ForgetPassword")
	for d := range msgs {
		log.Println("[FP - x] %s", d.Body)

		err := json.Unmarshal(d.Body, &fp)
		if err != nil {
			panic(err)
		}

		log.Println("[Data] %s", fp)
	}
}
