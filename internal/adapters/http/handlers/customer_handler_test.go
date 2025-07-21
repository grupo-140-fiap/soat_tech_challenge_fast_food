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
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) CreateCustomer(customer *dto.CreateCustomerDTO) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerService) GetCustomerByCpf(cpf string) (*entities.Customer, error) {
	args := m.Called(cpf)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func TestNewCustomerHandler(t *testing.T) {
	mockService := &MockCustomerService{}
	handler := NewCustomerHandler(mockService)

	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.customerService)
}

func TestCustomerHandler_CreateCustomer_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	customerDTO := dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao.silva@email.com",
	}

	mockService := &MockCustomerService{}
	mockService.On("CreateCustomer", &customerDTO).Return(nil)

	handler := NewCustomerHandler(mockService)

	// Prepare request body
	requestBody, _ := json.Marshal(customerDTO)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/customers", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.CreateCustomer(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Customer created successfully", response["message"])

	mockService.AssertExpectations(t)
}

func TestCustomerHandler_CreateCustomer_InvalidJSON(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	mockService := &MockCustomerService{}
	handler := NewCustomerHandler(mockService)

	// Invalid JSON body
	invalidJSON := `{"firstName": "João", "lastName": "Silva", "cpf":}`

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/customers", bytes.NewBuffer([]byte(invalidJSON)))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.CreateCustomer(c)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid input", response["message"])

	mockService.AssertNotCalled(t, "CreateCustomer")
}

func TestCustomerHandler_CreateCustomer_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	customerDTO := dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao.silva@email.com",
	}

	expectedError := errors.New("database connection failed")

	mockService := &MockCustomerService{}
	mockService.On("CreateCustomer", &customerDTO).Return(expectedError)

	handler := NewCustomerHandler(mockService)

	// Prepare request body
	requestBody, _ := json.Marshal(customerDTO)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/customers", bytes.NewBuffer(requestBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Act
	handler.CreateCustomer(c)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to create customer", response["message"])
	assert.Equal(t, "database connection failed", response["error"])

	mockService.AssertExpectations(t)
}

func TestCustomerHandler_GetCustomerByCpf_Success(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	cpf := "123.456.789-00"
	mockCustomer := &entities.Customer{
		ID:        1,
		FirstName: "João",
		LastName:  "Silva",
		CPF:       cpf,
		Email:     "joao.silva@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockService := &MockCustomerService{}
	mockService.On("GetCustomerByCpf", cpf).Return(mockCustomer, nil)

	handler := NewCustomerHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/customers/"+cpf, nil)
	c.Params = gin.Params{
		{Key: "cpf", Value: cpf},
	}

	// Act
	handler.GetCustomerByCpf(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var responseCustomer entities.Customer
	err := json.Unmarshal(w.Body.Bytes(), &responseCustomer)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), responseCustomer.ID)
	assert.Equal(t, "João", responseCustomer.FirstName)
	assert.Equal(t, "Silva", responseCustomer.LastName)
	assert.Equal(t, cpf, responseCustomer.CPF)
	assert.Equal(t, "joao.silva@email.com", responseCustomer.Email)

	mockService.AssertExpectations(t)
}

func TestCustomerHandler_GetCustomerByCpf_ServiceError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)

	cpf := "123.456.789-00"
	expectedError := errors.New("customer not found")

	mockService := &MockCustomerService{}
	mockService.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	handler := NewCustomerHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/customers/"+cpf, nil)
	c.Params = gin.Params{
		{Key: "cpf", Value: cpf},
	}

	// Act
	handler.GetCustomerByCpf(c)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed on find customer", response["message"])
	assert.Equal(t, "customer not found", response["error"])

	mockService.AssertExpectations(t)
}

func TestCustomerHandler_GetCustomerByCpf_EmptyCpf(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cpf := ""
	expectedError := errors.New("invalid CPF")

	mockService := &MockCustomerService{}
	mockService.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	handler := NewCustomerHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/customers/", nil)
	c.Params = gin.Params{
		{Key: "cpf", Value: cpf},
	}

	handler.GetCustomerByCpf(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed on find customer", response["message"])
	assert.Equal(t, "invalid CPF", response["error"])

	mockService.AssertExpectations(t)
}
