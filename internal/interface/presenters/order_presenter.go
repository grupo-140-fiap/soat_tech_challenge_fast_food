package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

// OrderPresenter handles the presentation logic for order responses
type OrderPresenter interface {
	PresentOrder(order *dto.OrderResponse) interface{}
	PresentOrders(orders []*dto.OrderResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

// orderPresenter implements OrderPresenter interface
type orderPresenter struct{}

// NewOrderPresenter creates a new order presenter
func NewOrderPresenter() OrderPresenter {
	return &orderPresenter{}
}

// PresentOrder formats a single order response
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

// PresentOrders formats multiple order responses
func (p *orderPresenter) PresentOrders(orders []*dto.OrderResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Orders retrieved successfully",
		"data":    orders,
		"count":   len(orders),
	}
}

// PresentError formats error responses
func (p *orderPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
}

// PresentSuccess formats success responses
func (p *orderPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    nil,
	}
}
