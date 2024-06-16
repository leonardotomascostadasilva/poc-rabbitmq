package config

import (
	producer "poc-rabbitmq/cmd/api/publish/Producer"
)

type App struct {
	Producer *producer.Producer
}

func InitializeApp() (*App, error) {
	conn, err := Connect()
	if err != nil {
		return nil, err
	}

	producer, err := producer.NewProducer(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &App{
		Producer: producer,
	}, nil
}
