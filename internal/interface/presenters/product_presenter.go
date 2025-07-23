package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

// ProductPresenter handles the presentation logic for product responses
type ProductPresenter interface {
	PresentProduct(product *dto.ProductResponse) interface{}
	PresentProducts(products []*dto.ProductResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

// productPresenter implements ProductPresenter interface
type productPresenter struct{}

// NewProductPresenter creates a new product presenter
func NewProductPresenter() ProductPresenter {
	return &productPresenter{}
}

// PresentProduct formats a single product response
func (p *productPresenter) PresentProduct(product *dto.ProductResponse) interface{} {
	if product == nil {
		return map[string]interface{}{
			"success": false,
			"message": "Product not found",
			"data":    nil,
		}
	}

	return map[string]interface{}{
		"success": true,
		"message": "Product retrieved successfully",
		"data":    product,
	}
}

// PresentProducts formats multiple product responses
func (p *productPresenter) PresentProducts(products []*dto.ProductResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Products retrieved successfully",
		"data":    products,
		"count":   len(products),
	}
}

// PresentError formats error responses
func (p *productPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
}

// PresentSuccess formats success responses
func (p *productPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    nil,
	}
}
