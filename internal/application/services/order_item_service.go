package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type OrderItemService struct {
	orderItemRepository repositories.OrderItemRepository
}

func NewOrderItemService(orderItemRepository repositories.OrderItemRepository) *OrderItemService {
	return &OrderItemService{
		orderItemRepository: orderItemRepository,
	}
}

func (u *OrderItemService) CreateOrderItem(order *dto.OrderItemDTO) error {
	err := u.orderItemRepository.CreateOrderItem(order)

	if err != nil {
		return err
	}

	return nil
}
