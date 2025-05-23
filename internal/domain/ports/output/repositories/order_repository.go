package repositories

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type OrderRepository interface {
	CreateOrder(order *dto.OrderDTO) error
}
