package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProductById(id string) (*entities.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Product), args.Error(1)
}

func (m *MockProductService) GetProductByCategory(category string) ([]entities.Product, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockProductService) CreateProduct(product *dto.ProductDTO) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) UpdateProduct(product *dto.ProductDTO) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) DeleteProductById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestNewProductHandler(t *testing.T) {
	mockService := &MockProductService{}
	handler := NewProductHandler(mockService)

	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.productService)
}

func TestProductHandler_GetProductById_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productID := "1"
	mockProduct := &entities.Product{
		ID:          1,
		Name:        "Burger",
		Description: "Delicious burger",
		Price:       15.99,
		Category:    "main",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockService := &MockProductService{}
	mockService.On("GetProductById", productID).Return(mockProduct, nil)

	handler := NewProductHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/products/1", nil)
	c.Params = gin.Params{
		{Key: "id", Value: productID},
	}

	handler.GetProductById(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseProduct entities.Product
	err := json.Unmarshal(w.Body.Bytes(), &responseProduct)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), responseProduct.ID)
	assert.Equal(t, "Burger", responseProduct.Name)
	assert.Equal(t, "main", responseProduct.Category)

	mockService.AssertExpectations(t)
}

func TestProductHandler_GetProductById_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productID := "1"
	expectedError := errors.New("product not found")

	mockService := &MockProductService{}
	mockService.On("GetProductById", productID).Return(nil, expectedError)

	handler := NewProductHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/products/1", nil)
	c.Params = gin.Params{
		{Key: "id", Value: productID},
	}

	handler.GetProductById(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed on find product", response["message"])
	assert.Equal(t, "product not found", response["error"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_GetProductByCategory_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	category := "main"
	mockProducts := []entities.Product{
		{
			ID:          1,
			Name:        "Burger",
			Description: "Delicious burger",
			Price:       15.99,
			Category:    "main",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Pizza",
			Description: "Tasty pizza",
			Price:       25.50,
			Category:    "main",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	mockService := &MockProductService{}
	mockService.On("GetProductByCategory", category).Return(mockProducts, nil)

	handler := NewProductHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/products/category/main", nil)
	c.Params = gin.Params{
		{Key: "category", Value: category},
	}

	handler.GetProductByCategory(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseProducts []entities.Product
	err := json.Unmarshal(w.Body.Bytes(), &responseProducts)
	assert.NoError(t, err)
	assert.Len(t, responseProducts, 2)
	assert.Equal(t, "Burger", responseProducts[0].Name)
	assert.Equal(t, "Pizza", responseProducts[1].Name)

	mockService.AssertExpectations(t)
}

func TestProductHandler_GetProductByCategory_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	category := "main"
	expectedError := errors.New("database connection failed")

	mockService := &MockProductService{}
	mockService.On("GetProductByCategory", category).Return(([]entities.Product)(nil), expectedError)

	handler := NewProductHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/products/category/main", nil)
	c.Params = gin.Params{
		{Key: "category", Value: category},
	}

	handler.GetProductByCategory(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed on find product", response["message"])
	assert.Equal(t, "database connection failed", response["error"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_CreateProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productDTO := dto.ProductDTO{
		Name:        "New Burger",
		Description: "A new delicious burger",
		Price:       "18.99",
		Category:    "main",
	}

	mockService := &MockProductService{}
	mockService.On("CreateProduct", &productDTO).Return(nil)

	handler := NewProductHandler(mockService)

	requestBody, _ := json.Marshal(productDTO)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateProduct(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Product created successfully", response["message"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_CreateProduct_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockProductService{}
	handler := NewProductHandler(mockService)

	invalidJSON := `{"name": "Burger", "price":}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateProduct(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid input", response["message"])

	mockService.AssertNotCalled(t, "CreateProduct")
}

func TestProductHandler_CreateProduct_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productDTO := dto.ProductDTO{
		Name:        "New Burger",
		Description: "A new delicious burger",
		Price:       "18.99",
		Category:    "main",
	}

	expectedError := errors.New("failed to save product")

	mockService := &MockProductService{}
	mockService.On("CreateProduct", &productDTO).Return(expectedError)

	handler := NewProductHandler(mockService)

	requestBody, _ := json.Marshal(productDTO)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateProduct(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to create product", response["message"])
	assert.Equal(t, "failed to save product", response["error"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_UpdateProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productDTO := dto.ProductDTO{
		ID:          1,
		Name:        "Updated Burger",
		Description: "An updated delicious burger",
		Price:       "19.99",
		Category:    "main",
	}

	mockService := &MockProductService{}
	mockService.On("UpdateProduct", &productDTO).Return(nil)

	handler := NewProductHandler(mockService)

	requestBody, _ := json.Marshal(productDTO)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/products", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateProduct(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Product updated successfully", response["message"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_UpdateProduct_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockProductService{}
	handler := NewProductHandler(mockService)

	invalidJSON := `{"id": 1, "name": "Burger", "price":}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/products", bytes.NewBuffer([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateProduct(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid input", response["message"])

	mockService.AssertNotCalled(t, "UpdateProduct")
}

func TestProductHandler_UpdateProduct_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productDTO := dto.ProductDTO{
		ID:          1,
		Name:        "Updated Burger",
		Description: "An updated delicious burger",
		Price:       "19.99",
		Category:    "main",
	}

	expectedError := errors.New("failed to update product")

	mockService := &MockProductService{}
	mockService.On("UpdateProduct", &productDTO).Return(expectedError)

	handler := NewProductHandler(mockService)

	requestBody, _ := json.Marshal(productDTO)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPut, "/products", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateProduct(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to update product", response["message"])
	assert.Equal(t, "failed to update product", response["error"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_DeleteProductById_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productID := "1"

	mockService := &MockProductService{}
	mockService.On("DeleteProductById", productID).Return(nil)

	handler := NewProductHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	c.Params = gin.Params{
		{Key: "id", Value: productID},
	}

	handler.DeleteProductById(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Product deleted successfully", response["message"])

	mockService.AssertExpectations(t)
}

func TestProductHandler_DeleteProductById_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	productID := "1"
	expectedError := errors.New("failed to delete product")

	mockService := &MockProductService{}
	mockService.On("DeleteProductById", productID).Return(expectedError)

	handler := NewProductHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	c.Params = gin.Params{
		{Key: "id", Value: productID},
	}

	handler.DeleteProductById(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed on find product", response["message"])
	assert.Equal(t, "failed to delete product", response["error"])

	mockService.AssertExpectations(t)
}
