package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type OrderService interface {
	GetOrders() ([]entities.Order, error)
	CreateOrder(order *dto.OrderDTO) error
	GetOrderById(id string) (*entities.Order, error)
	UpdateOrderStatus(id string, status string) error
}
