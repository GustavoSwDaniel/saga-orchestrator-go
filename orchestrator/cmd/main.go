package main

import (
	"log"
	"orchestrator/internal/infrastructure/rabbitmq"
	"orchestrator/internal/usecase/createorder"

	"orchestrator/pkg/config"
)

func main() {
	configs := config.LoadConfig()
	conn, err := rabbitmq.Connect(configs.RabbitMqUrl)
	if err != nil {
		log.Fatalf("Error to connect in brocker: %v", err)
	}
	defer conn.Close()

	producer := rabbitmq.NewProducer(conn)

	producerService := createorder.NewCreateOrderUseCase(producer)

	consumer := rabbitmq.NewConsumer(conn)
	if err := consumer.Consumer("events_saga.queue", producerService.Execute); err != nil {
		log.Fatal(err)
	}
}
