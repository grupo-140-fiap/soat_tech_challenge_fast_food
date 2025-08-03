package output

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"

// ProductGateway defines the contract for product data access operations
type ProductGateway interface {
	Create(product *entities.Product) error
	GetByID(id uint64) (*entities.Product, error)
	GetByCategory(category string) ([]*entities.Product, error)
	GetAll() ([]*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id uint64) error
}
