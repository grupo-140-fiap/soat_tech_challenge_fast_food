package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type OrderPresenter interface {
	PresentOrder(order *dto.OrderResponse) interface{}
	PresentOrders(orders []*dto.OrderResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

type orderPresenter struct{}

func NewOrderPresenter() OrderPresenter {
	return &orderPresenter{}
}

func (p *orderPresenter) PresentOrder(order *dto.OrderResponse) interface{} {
	if order == nil {
		return map[string]interface{}{
			"success": false,
			"message": "Order not found",
			"data":    nil,
		}
	}

	return map[string]interface{}{
		"success": true,
		"message": "Order retrieved successfully",
		"data":    order,
	}
}

func (p *orderPresenter) PresentOrders(orders []*dto.OrderResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Orders retrieved successfully",
		"data":    orders,
		"count":   len(orders),
	}
}

func (p *orderPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
}

func (p *orderPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    nil,
	}
}
