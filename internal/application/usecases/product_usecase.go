package usecases

import (
	"errors"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/repositories"
)

// ProductUseCase defines the interface for product business operations
// Following Interface Segregation Principle
type ProductUseCase interface {
	CreateProduct(request *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetProductByID(id uint64) (*dto.ProductResponse, error)
	GetProductsByCategory(category string) ([]*dto.ProductResponse, error)
	GetAllProducts() ([]*dto.ProductResponse, error)
	UpdateProduct(id uint64, request *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(id uint64) error
}

// productUseCase implements ProductUseCase interface
type productUseCase struct {
	productRepo repositories.ProductRepository
}

// NewProductUseCase creates a new product use case
func NewProductUseCase(productRepo repositories.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product
func (uc *productUseCase) CreateProduct(request *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	// Business validation for category
	if !entities.IsValidCategory(request.Category) {
		return nil, errors.New("invalid product category")
	}

	// Create domain entity
	product := entities.NewProduct(
		request.Name,
		request.Description,
		request.Price,
		entities.ProductCategory(request.Category),
		request.ImageUrl,
	)

	// Business validation
	if !product.IsValid() {
		return nil, errors.New("invalid product data")
	}

	// Persist entity
	err := uc.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	// Return response DTO
	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		ImageUrl:    product.ImageUrl,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// GetProductByID retrieves a product by ID
func (uc *productUseCase) GetProductByID(id uint64) (*dto.ProductResponse, error) {
	product, err := uc.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		ImageUrl:    product.ImageUrl,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// GetProductsByCategory retrieves products by category
func (uc *productUseCase) GetProductsByCategory(category string) ([]*dto.ProductResponse, error) {
	// Business validation for category
	if !entities.IsValidCategory(category) {
		return nil, errors.New("invalid product category")
	}

	products, err := uc.productRepo.GetByCategory(category)
	if err != nil {
		return nil, err
	}

	var response []*dto.ProductResponse
	for _, product := range products {
		response = append(response, &dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Category:    product.Category,
			ImageUrl:    product.ImageUrl,
			CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return response, nil
}

// GetAllProducts retrieves all products
func (uc *productUseCase) GetAllProducts() ([]*dto.ProductResponse, error) {
	products, err := uc.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []*dto.ProductResponse
	for _, product := range products {
		response = append(response, &dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Category:    product.Category,
			ImageUrl:    product.ImageUrl,
			CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return response, nil
}

// UpdateProduct updates an existing product
func (uc *productUseCase) UpdateProduct(id uint64, request *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	// Business validation for category
	if !entities.IsValidCategory(request.Category) {
		return nil, errors.New("invalid product category")
	}

	product, err := uc.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	// Update entity using domain method
	product.UpdateProduct(
		request.Name,
		request.Description,
		request.Price,
		entities.ProductCategory(request.Category),
		request.ImageUrl,
	)

	// Business validation
	if !product.IsValid() {
		return nil, errors.New("invalid product data")
	}

	// Persist changes
	err = uc.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		ImageUrl:    product.ImageUrl,
		CreatedAt:   product.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// DeleteProduct deletes a product
func (uc *productUseCase) DeleteProduct(id uint64) error {
	product, err := uc.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found")
	}

	return uc.productRepo.Delete(id)
}
