package rabbitmq

type IProducer interface {
	Produce(exchange string, routingKey string, message []byte) error
}
