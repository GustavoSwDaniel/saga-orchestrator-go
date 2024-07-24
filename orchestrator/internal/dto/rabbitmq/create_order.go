package rabbitmq

type OrderMessage struct {
	ProductUUID string `json:"product_uuid"`
	Quantity    int    `json:"quantity"`
}

type CreateOrder struct {
	EventType string       `json:"event_type"`
	Message   OrderMessage `json:"message"`
}

type SendMenssage struct {
	EventType     string       `json:"event_type"`
	Message       OrderMessage `json:"message"`
	TransactionId string       `json:"transaction_id"`
}
