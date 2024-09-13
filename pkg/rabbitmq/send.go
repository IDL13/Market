package rabbitmq

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer interface {
	WriteData(body []byte) int
}

func New(name string) rabbitMQProducer {

	conn := connect()
	ch := createRabbitMQChannel(conn)

	createQueue(ch, name)

	return rabbitMQProducer{
		conn: conn,
		ch:   ch,
		name: name,
	}
}

type rabbitMQProducer struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	name string
}

func (r rabbitMQProducer) WriteData(body []byte) int {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer r.conn.Close()
	defer r.ch.Close()
	defer cancel()

	err := r.ch.PublishWithContext(ctx,
		"",
		r.name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		errorChecker(err, "Failed to publish a message")
		return 0
	} else {
		return 1
	}
}

func errorChecker(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func connect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errorChecker(err, "Failed to connect RabbitMQ")

	return conn
}

func createRabbitMQChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	errorChecker(err, "Failed to open a channel")

	return ch
}

func createQueue(ch *amqp.Channel, name string) {
	_, err := ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)

	errorChecker(err, "Failed to declare a queue")
}
