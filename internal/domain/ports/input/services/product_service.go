package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type ProductService interface {
	CreateProduct(product *dto.CreateProductDTO) error
	GetProductById(cpf string) (*entities.Product, error)
}
