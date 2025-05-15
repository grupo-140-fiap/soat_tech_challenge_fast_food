package repositories

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type ProductRepository interface {
	CreateProduct(product *dto.CreateProductDTO) error
	GetProductById(id string) (*entities.Product, error)
}
