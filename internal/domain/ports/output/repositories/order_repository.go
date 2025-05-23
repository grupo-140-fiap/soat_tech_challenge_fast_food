package repositories

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type OrderRepository interface {
	CreateOrder(order *dto.OrderDTO) error
	GetOrderById(id string) (*entities.Orders, error)
}
