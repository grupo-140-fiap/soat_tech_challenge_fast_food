package usecases

import (
	"errors"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/repositories"
)

type ProductUseCase interface {
	CreateProduct(request *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetProductByID(id uint64) (*dto.ProductResponse, error)
	GetProductsByCategory(category string) ([]*dto.ProductResponse, error)
	GetAllProducts() ([]*dto.ProductResponse, error)
	UpdateProduct(id uint64, request *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(id uint64) error
}

type productUseCase struct {
	productRepo repositories.ProductRepository
}

func NewProductUseCase(productRepo repositories.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepo: productRepo,
	}
}

func (uc *productUseCase) CreateProduct(request *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	if !entities.IsValidCategory(request.Category) {
		return nil, errors.New("invalid product category")
	}

	product := entities.NewProduct(
		request.Name,
		request.Description,
		request.Price,
		entities.ProductCategory(request.Category),
		request.ImageUrl,
	)

	if !product.IsValid() {
		return nil, errors.New("invalid product data")
	}

	err := uc.productRepo.Create(product)
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

func (uc *productUseCase) GetProductsByCategory(category string) ([]*dto.ProductResponse, error) {
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

func (uc *productUseCase) UpdateProduct(id uint64, request *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
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

	product.UpdateProduct(
		request.Name,
		request.Description,
		request.Price,
		entities.ProductCategory(request.Category),
		request.ImageUrl,
	)

	if !product.IsValid() {
		return nil, errors.New("invalid product data")
	}

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
