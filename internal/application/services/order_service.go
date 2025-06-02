package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type OrderService struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (u *OrderService) GetOrders() ([]entities.Order, error) {
	orders, err := u.orderRepository.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (u *OrderService) CreateOrder(order *dto.OrderDTO) error {
	return u.orderRepository.CreateOrder(order)
}

func (u *OrderService) GetOrderById(id string) (*entities.Order, error) {
	order, err := u.orderRepository.GetOrderById(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (u *OrderService) UpdateOrderStatus(id string, status string) error {
	return u.orderRepository.UpdateOrderStatus(id, status)
}
