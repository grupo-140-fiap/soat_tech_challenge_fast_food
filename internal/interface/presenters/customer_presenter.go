package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type CustomerPresenter interface {
	PresentCustomer(customer *dto.CustomerResponse) interface{}
	PresentCustomers(customers []*dto.CustomerResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

type customerPresenter struct{}

func NewCustomerPresenter() CustomerPresenter {
	return &customerPresenter{}
}

func (p *customerPresenter) PresentCustomer(customer *dto.CustomerResponse) interface{} {
	if customer == nil {
		return map[string]interface{}{
			"success": false,
			"message": "Customer not found",
			"data":    nil,
		}
	}

	return map[string]interface{}{
		"success": true,
		"message": "Customer retrieved successfully",
		"data":    customer,
	}
}

func (p *customerPresenter) PresentCustomers(customers []*dto.CustomerResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Customers retrieved successfully",
		"data":    customers,
		"count":   len(customers),
	}
}

func (p *customerPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
}

func (p *customerPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    nil,
	}
}
