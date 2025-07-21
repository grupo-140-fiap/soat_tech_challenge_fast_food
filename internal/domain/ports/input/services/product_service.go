package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
)

type ProductService interface {
	GetProductById(id string) (*entities.Product, error)
	GetProductByCategory(category string) ([]entities.Product, error)

	CreateProduct(product *dto.ProductDTO) error
	UpdateProduct(producId int, product *dto.ProductDTO) error

	DeleteProductById(id string) error
}
