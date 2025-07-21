package repositories

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
)

type OrderRepository interface {
	GetOrders() ([]entities.Order, error)
	CreateOrder(order *dto.OrderDTO) error
	GetOrderById(id string) (*entities.Order, error)
	UpdateOrderStatus(id string, status string) error
	GetActiveOrders() (*[]entities.Order, error)
}
