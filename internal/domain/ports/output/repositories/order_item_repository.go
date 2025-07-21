package repositories

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
)

type OrderItemRepository interface {
	CreateOrderItem(item *dto.OrderItemDTO) error
}
