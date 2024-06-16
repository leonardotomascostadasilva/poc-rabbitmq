package producer

import (
	"encoding/json"
	"fmt"
	domain "poc-rabbitmq/cmd/api/publish/Domain"

	"github.com/streadway/amqp"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewProducer(conn *amqp.Connection) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	q, err := ch.QueueDeclare(
		"golang-queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &Producer{
		conn:    conn,
		channel: ch,
		queue:   q,
	}, nil
}

func (p *Producer) Execute(msg *domain.Message) error {
	messageBody, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message to JSON: %w", err)
	}

	err = p.channel.Publish(
		"",           // exchange
		p.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	return nil
}

func (p *Producer) Close() {
	p.channel.Close()
	p.conn.Close()
}
