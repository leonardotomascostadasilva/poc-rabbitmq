package config

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}
