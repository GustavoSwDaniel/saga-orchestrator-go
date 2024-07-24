package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	conn *amqp.Connection
}

func NewProducer(conn *amqp.Connection) IProducer {
	return &Producer{conn: conn}
}

func (p *Producer) Produce(exchange string, routingKey string, message []byte) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err = ch.Publish(exchange, routingKey, false, false, amqp.Publishing{
		Body: []byte(message),
	}); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
