package usecase

import (
	"github.com/axnd/goaxnd/internal/order/entity"
	"github.com/axnd/goaxnd/internal/order/infra/database"
)

type OrderInputDTO struct {
	ID    int
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         int
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: &orderRepository}
}

// pega o dado da entrada e retorna o da saida
func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CheckFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDTO{ID: order.ID, Price: order.Price, Tax: order.Tax, FinalPrice: order.FinalPrice}, nil
}
