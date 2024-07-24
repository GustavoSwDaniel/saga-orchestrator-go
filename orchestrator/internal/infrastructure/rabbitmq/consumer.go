package rabbitmq

import (
	"encoding/json"
	"log"

	"orchestrator/internal/dto/rabbitmq"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn *amqp.Connection
}

func NewConsumer(conn *amqp.Connection) *Consumer {
	return &Consumer{conn: conn}
}

func (c *Consumer) Consumer(queueName string, execution func(*rabbitmq.CreateOrder)) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			data := &rabbitmq.CreateOrder{}
			err := json.Unmarshal(d.Body, data)
			if err != nil {
				log.Fatal("Error na deserialização")
			}
			log.Printf("Mensagem %+v", data)
			execution(data)
		}
	}()

	log.Printf("Esperando mesagem")
	<-forever
	return nil
}
