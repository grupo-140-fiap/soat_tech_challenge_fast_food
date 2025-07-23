package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

// CustomerPresenter handles the presentation logic for customer responses
// Following the Single Responsibility Principle
type CustomerPresenter interface {
	PresentCustomer(customer *dto.CustomerResponse) interface{}
	PresentCustomers(customers []*dto.CustomerResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

// customerPresenter implements CustomerPresenter interface
type customerPresenter struct{}

// NewCustomerPresenter creates a new customer presenter
func NewCustomerPresenter() CustomerPresenter {
	return &customerPresenter{}
}

// PresentCustomer formats a single customer response
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

// PresentCustomers formats multiple customer responses
func (p *customerPresenter) PresentCustomers(customers []*dto.CustomerResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Customers retrieved successfully",
		"data":    customers,
		"count":   len(customers),
	}
}

// PresentError formats error responses
func (p *customerPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
}

// PresentSuccess formats success responses
func (p *customerPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    nil,
	}
}
