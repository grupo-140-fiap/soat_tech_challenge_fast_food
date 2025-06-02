package services

import (
	"fmt"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

var validOrderStatuses = map[string]bool{
	"received":    true,
	"preparation": true,
	"ready":       true,
	"completed":   true,
}

type OrderService struct {
	orderRepository repositories.OrderRepository
	productService  services.ProductService
}

func NewOrderService(orderRepository repositories.OrderRepository, productService services.ProductService) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productService:  productService,
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
	for _, item := range order.Items {
		_, err := u.productService.GetProductById(fmt.Sprintf("%d", item.ProductId))
		if err != nil {
			return fmt.Errorf("failed to validate product with ID %d: %w", item.ProductId, err)
		}
	}

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
	if !isValidOrderStatus(status) {
		return fmt.Errorf("invalid order status: %s. Valid statuses are: received, preparation, ready, completed", status)
	}

	return u.orderRepository.UpdateOrderStatus(id, status)
}

func isValidOrderStatus(status string) bool {
	return validOrderStatuses[status]
}
