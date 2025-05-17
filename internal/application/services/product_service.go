package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type ProductService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (u *ProductService) CreateProduct(product *dto.ProductDTO) error {
	err := u.productRepository.CreateProduct(product)

	if err != nil {
		return err
	}

	return nil
}

func (u *ProductService) GetProductById(id string) (*entities.Product, error) {
	product, err := u.productRepository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (u *ProductService) UpdateProduct(product *dto.ProductDTO) error {
	err := u.productRepository.UpdateProduct(product)

	if err != nil {
		return err
	}

	return nil
}

func (u *ProductService) DeleteProductById(id string) error {
	err := u.productRepository.DeleteProductById(id)

	if err != nil {
		return err
	}

	return nil
}
