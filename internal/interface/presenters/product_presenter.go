package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type ProductPresenter interface {
	PresentProduct(product *dto.ProductResponse) interface{}
	PresentProducts(products []*dto.ProductResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

type productPresenter struct{}

func NewProductPresenter() ProductPresenter {
	return &productPresenter{}
}

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

func (p *productPresenter) PresentProducts(products []*dto.ProductResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "Products retrieved successfully",
		"data":    products,
		"count":   len(products),
	}
}

func (p *productPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
}

func (p *productPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    nil,
	}
}
