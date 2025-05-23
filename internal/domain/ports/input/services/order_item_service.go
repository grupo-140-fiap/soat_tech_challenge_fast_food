package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type OrderItemService interface {
	CreateOrderItem(item *dto.OrderItemDTO) error
}
