package services

import (
	"errors"
	"testing"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) CreateProduct(product *dto.ProductDTO) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetProductById(id string) (*entities.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByCategory(category string) ([]entities.Product, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductRepository) UpdateProduct(product *dto.ProductDTO) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProductById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewProductService(t *testing.T) {
	mockRepo := &MockProductRepository{}
	service := NewProductService(mockRepo)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.productRepository)
}

func TestProductService_CreateProduct_Success(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productDTO := &dto.ProductDTO{
		Name:        "Cheeseburger",
		Description: "Delicious cheeseburger",
		Price:       "19.90",
		Category:    "Lanche",
		Image:       "https://example.com/cheeseburger.png",
	}

	mockRepo.On("CreateProduct", productDTO).Return(nil)

	service := NewProductService(mockRepo)

	err := service.CreateProduct(productDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_CreateProduct_RepositoryError(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productDTO := &dto.ProductDTO{
		Name:        "Cheeseburger",
		Description: "Delicious cheeseburger",
		Price:       "19.90",
		Category:    "Lanche",
		Image:       "https://example.com/cheeseburger.png",
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("CreateProduct", productDTO).Return(expectedError)

	service := NewProductService(mockRepo)

	err := service.CreateProduct(productDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductById_Success(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productID := "1"
	expectedProduct := &entities.Product{
		ID:          1,
		Name:        "Cheeseburger",
		Description: "Delicious cheeseburger",
		Price:       19.90,
		Category:    "Lanche",
		Image:       "https://example.com/cheeseburger.png",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("GetProductById", productID).Return(expectedProduct, nil)

	service := NewProductService(mockRepo)

	product, err := service.GetProductById(productID)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, expectedProduct.ID, product.ID)
	assert.Equal(t, expectedProduct.Name, product.Name)
	assert.Equal(t, expectedProduct.Description, product.Description)
	assert.Equal(t, expectedProduct.Price, product.Price)
	assert.Equal(t, expectedProduct.Category, product.Category)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductById_NotFound(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productID := "999"
	expectedError := errors.New("product not found")

	mockRepo.On("GetProductById", productID).Return(nil, expectedError)

	service := NewProductService(mockRepo)

	product, err := service.GetProductById(productID)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductById_RepositoryError(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productID := "1"
	expectedError := errors.New("database connection failed")

	mockRepo.On("GetProductById", productID).Return(nil, expectedError)

	service := NewProductService(mockRepo)

	product, err := service.GetProductById(productID)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductByCategory_Success(t *testing.T) {
	mockRepo := &MockProductRepository{}

	category := "Lanche"
	expectedProducts := []entities.Product{
		{
			ID:          1,
			Name:        "Cheeseburger",
			Description: "Delicious cheeseburger",
			Price:       19.90,
			Category:    "Lanche",
			Image:       "https://example.com/cheeseburger.png",
		},
		{
			ID:          2,
			Name:        "Big Burger",
			Description: "Big delicious burger",
			Price:       25.90,
			Category:    "Lanche",
			Image:       "https://example.com/bigburger.png",
		},
	}

	mockRepo.On("GetProductByCategory", category).Return(expectedProducts, nil)

	service := NewProductService(mockRepo)

	products, err := service.GetProductByCategory(category)

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 2)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[1].Name, products[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductByCategory_EmptyResult(t *testing.T) {
	mockRepo := &MockProductRepository{}

	category := "nonexistent"
	expectedProducts := []entities.Product{}

	mockRepo.On("GetProductByCategory", category).Return(expectedProducts, nil)

	service := NewProductService(mockRepo)

	products, err := service.GetProductByCategory(category)

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 0)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetProductByCategory_RepositoryError(t *testing.T) {
	mockRepo := &MockProductRepository{}

	category := "Lanche"
	expectedError := errors.New("database connection failed")

	mockRepo.On("GetProductByCategory", category).Return(nil, expectedError)

	service := NewProductService(mockRepo)

	products, err := service.GetProductByCategory(category)

	assert.Error(t, err)
	assert.Nil(t, products)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_Success(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productDTO := &dto.ProductDTO{
		ID:          1,
		Name:        "Updated Cheeseburger",
		Description: "Updated delicious cheeseburger",
		Price:       "21.90",
		Category:    "Lanche",
		Image:       "https://example.com/updated-cheeseburger.png",
	}

	mockRepo.On("UpdateProduct", productDTO).Return(nil)

	service := NewProductService(mockRepo)

	err := service.UpdateProduct(productDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_RepositoryError(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productDTO := &dto.ProductDTO{
		ID:          1,
		Name:        "Updated Cheeseburger",
		Description: "Updated delicious cheeseburger",
		Price:       "21.90",
		Category:    "Lanche",
		Image:       "https://example.com/updated-cheeseburger.png",
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("UpdateProduct", productDTO).Return(expectedError)

	service := NewProductService(mockRepo)

	err := service.UpdateProduct(productDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProductById_Success(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productID := "1"

	mockRepo.On("DeleteProductById", productID).Return(nil)

	service := NewProductService(mockRepo)

	err := service.DeleteProductById(productID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProductById_RepositoryError(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productID := "1"
	expectedError := errors.New("database connection failed")

	mockRepo.On("DeleteProductById", productID).Return(expectedError)

	service := NewProductService(mockRepo)

	err := service.DeleteProductById(productID)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProductById_NotFound(t *testing.T) {
	mockRepo := &MockProductRepository{}

	productID := "999"
	expectedError := errors.New("product not found")

	mockRepo.On("DeleteProductById", productID).Return(expectedError)

	service := NewProductService(mockRepo)

	err := service.DeleteProductById(productID)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_CreateProduct_NilProduct(t *testing.T) {
	mockRepo := &MockProductRepository{}

	expectedError := errors.New("invalid product data")
	mockRepo.On("CreateProduct", (*dto.ProductDTO)(nil)).Return(expectedError)

	service := NewProductService(mockRepo)

	err := service.CreateProduct(nil)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_NilProduct(t *testing.T) {
	mockRepo := &MockProductRepository{}

	expectedError := errors.New("invalid product data")
	mockRepo.On("UpdateProduct", (*dto.ProductDTO)(nil)).Return(expectedError)

	service := NewProductService(mockRepo)

	err := service.UpdateProduct(nil)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
