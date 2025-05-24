package repositories

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type OrderItemRepository interface {
	CreateOrderItem(item *dto.OrderItemDTO) error
}
