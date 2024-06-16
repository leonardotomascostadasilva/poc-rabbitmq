package usecase

import (
	domain "poc-rabbitmq/cmd/api/publish/Domain"
	producer "poc-rabbitmq/cmd/api/publish/Producer"
)

type ProducerMessageUseCase struct {
	producer *producer.Producer
}

func NewProducerMessageUseCase(producer *producer.Producer) *ProducerMessageUseCase {
	return &ProducerMessageUseCase{
		producer: producer,
	}
}

func (uc *ProducerMessageUseCase) Execute(description, name string, status bool) error {
	message := domain.NewMessage(description, name, status)
	return uc.producer.Execute(message)
}
