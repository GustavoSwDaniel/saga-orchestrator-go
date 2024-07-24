package createorder

import dto "orchestrator/internal/dto/rabbitmq"

type ICreateOrderUseCase interface {
	Execute(body *dto.CreateOrder)
}
