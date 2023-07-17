package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	conn *amqp.Connection
}

func NewRabbitMQConsumer(conn *amqp.Connection) RabbitMQConsumer {
	return *&RabbitMQConsumer{
		conn: conn,
	}
}

func (m *RabbitMQConsumer) On(queue string, f func(deliveries <-chan amqp.Delivery)) {
	ch, err := m.conn.Channel()

	if err != nil {
		log.Fatalf("%s: %s", "Error al crear canal", err)
	}

	message, chErr := ch.Consume(
		queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if chErr != nil {
		log.Fatalf("%s: %s", "Error al consumir", chErr)
	}

	go f(message)

}
