package massange

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//SendMessange ,,,
func SendMessange() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"ok",  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	data := Message{
		ID:              1,
		ExternalID:      "2",
		Dst:             "2",
		Message:         "hi",
		Src:             "okkk",
		State:           1,
		CreatedAt:       &time.Time{},
		LastUpdatedDate: &time.Time{},
		SMSCMessageID:   "2",
	}
	body, _ := getByte(data)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func getByte(msg Message) ([]byte, error) {
	encoder, err := json.Marshal(&msg)
	return encoder, err
}
