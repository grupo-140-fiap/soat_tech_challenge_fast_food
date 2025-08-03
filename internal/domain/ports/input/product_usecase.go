package input

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"

// ProductUseCase defines the contract for product business operations
type ProductUseCase interface {
	CreateProduct(request *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetProductByID(id uint64) (*dto.ProductResponse, error)
	GetProductsByCategory(category string) ([]*dto.ProductResponse, error)
	GetAllProducts() ([]*dto.ProductResponse, error)
	UpdateProduct(id uint64, request *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(id uint64) error
}
