package main

import (
	"log"
	config "poc-rabbitmq/cmd/api/publish/Config"
	usecase "poc-rabbitmq/cmd/api/publish/UseCase"
)

func main() {

	app, err := config.InitializeApp()
	config.FailOnError(err, "Failed to initialize the application")
	defer app.Producer.Close()

	useCase := usecase.NewProducerMessageUseCase(app.Producer)

	err = useCase.Execute("This is a test message", "TestMessage", true)
	config.FailOnError(err, "Failed to publish a message")

	log.Println("Message published successfully")

}
