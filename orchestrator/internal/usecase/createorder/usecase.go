package createorder

import (
	"encoding/json"
	dto "orchestrator/internal/dto/rabbitmq"
	"orchestrator/internal/infrastructure/rabbitmq"

	"github.com/google/uuid"
)

type createOrderUseCase struct {
	BrokerProducer rabbitmq.IProducer
}

func NewCreateOrderUseCase(producer rabbitmq.IProducer) ICreateOrderUseCase {
	return &createOrderUseCase{
		BrokerProducer: producer,
	}
}

func (sp *createOrderUseCase) Execute(body *dto.CreateOrder) {
	transanctionId := uuid.New().String()
	sendMessage := dto.SendMenssage{
		EventType:     body.EventType,
		Message:       body.Message,
		TransactionId: transanctionId,
	}
	msg, err := json.Marshal(sendMessage)
	if err != nil {
		panic(err)
	}
	sp.BrokerProducer.Produce("EventsProcess", "invetory.queue", msg)
	sp.BrokerProducer.Produce("EventsProcess", "payment.queue", msg)
	sp.BrokerProducer.Produce("EventsProcess", "shipping.queue", msg)

}
