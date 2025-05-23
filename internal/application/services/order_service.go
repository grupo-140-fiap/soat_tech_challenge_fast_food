package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type OrderService struct {
	orderRepository     repositories.OrderRepository
	orderItemRepository repositories.OrderItemRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, orderItemRepository repositories.OrderItemRepository) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
	}
}

func (u *OrderService) CreateOrder(order *dto.OrderDTO) error {
	err := u.orderRepository.CreateOrder(order)

	if err != nil {
		return err
	} else {
		for _, item := range order.Items {
			err = u.orderItemRepository.CreateOrderItem(&item)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (u *OrderService) GetOrderById(id string) (*entities.Orders, error) {
	order, err := u.orderRepository.GetOrderById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *OrderService) UpdateOrderStatus(id string, status string) error {
	err := u.orderRepository.UpdateOrderStatus(id, status)

	if err != nil {
		return err
	}

	return nil
}
