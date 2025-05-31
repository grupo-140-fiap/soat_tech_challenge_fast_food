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

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) CreateCustomer(customer *dto.CreateCustomerDTO) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) GetCustomerByCpf(cpf string) (*entities.Customer, error) {
	args := m.Called(cpf)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Customer), args.Error(1)
}

func TestNewCustomerService(t *testing.T) {
	mockRepo := &MockCustomerRepository{}
	service := NewCustomerService(mockRepo)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.customerRepository)
}

func TestCustomerService_CreateCustomer_Success(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		Email:     "joao.silva@email.com",
		CPF:       "123.456.789-00",
	}

	mockRepo.On("CreateCustomer", customerDTO).Return(nil)

	service := NewCustomerService(mockRepo)

	err := service.CreateCustomer(customerDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_CreateCustomer_RepositoryError(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		Email:     "joao.silva@email.com",
		CPF:       "123.456.789-00",
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("CreateCustomer", customerDTO).Return(expectedError)

	service := NewCustomerService(mockRepo)

	err := service.CreateCustomer(customerDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_CreateCustomer_DuplicateCPF(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		Email:     "joao.silva@email.com",
		CPF:       "123.456.789-00",
	}

	expectedError := errors.New("UNIQUE constraint failed: customers.cpf")
	mockRepo.On("CreateCustomer", customerDTO).Return(expectedError)

	service := NewCustomerService(mockRepo)

	err := service.CreateCustomer(customerDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE constraint failed")
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_CreateCustomer_NilCustomer(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	expectedError := errors.New("invalid customer data")
	mockRepo.On("CreateCustomer", (*dto.CreateCustomerDTO)(nil)).Return(expectedError)

	service := NewCustomerService(mockRepo)

	err := service.CreateCustomer(nil)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_GetCustomerByCpf_Success(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	cpf := "123.456.789-00"
	expectedCustomer := &entities.Customer{
		ID:        1,
		FirstName: "João",
		LastName:  "Silva",
		CPF:       cpf,
		Email:     "joao.silva@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetCustomerByCpf", cpf).Return(expectedCustomer, nil)

	service := NewCustomerService(mockRepo)

	customer, err := service.GetCustomerByCpf(cpf)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, expectedCustomer.ID, customer.ID)
	assert.Equal(t, expectedCustomer.FirstName, customer.FirstName)
	assert.Equal(t, expectedCustomer.LastName, customer.LastName)
	assert.Equal(t, expectedCustomer.CPF, customer.CPF)
	assert.Equal(t, expectedCustomer.Email, customer.Email)
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_GetCustomerByCpf_NotFound(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	cpf := "999.999.999-99"
	expectedError := errors.New("customer with CPF 999.999.999-99 not found")

	mockRepo.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	service := NewCustomerService(mockRepo)

	customer, err := service.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Contains(t, err.Error(), "customer with CPF 999.999.999-99 not found")
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_GetCustomerByCpf_RepositoryError(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	cpf := "123.456.789-00"
	expectedError := errors.New("database connection failed")

	mockRepo.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	service := NewCustomerService(mockRepo)

	customer, err := service.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_GetCustomerByCpf_EmptyCPF(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	cpf := ""
	expectedError := errors.New("CPF cannot be empty")

	mockRepo.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	service := NewCustomerService(mockRepo)

	customer, err := service.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Contains(t, err.Error(), "CPF cannot be empty")
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_GetCustomerByCpf_InvalidCPFFormat(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	cpf := "invalid-cpf"
	expectedError := errors.New("invalid CPF format")

	mockRepo.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	service := NewCustomerService(mockRepo)

	customer, err := service.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Contains(t, err.Error(), "invalid CPF format")
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_CreateCustomer_DuplicateEmail(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "Maria",
		LastName:  "Santos",
		Email:     "joao.silva@email.com",
		CPF:       "987.654.321-00",
	}

	expectedError := errors.New("UNIQUE constraint failed: customers.email")
	mockRepo.On("CreateCustomer", customerDTO).Return(expectedError)

	service := NewCustomerService(mockRepo)

	err := service.CreateCustomer(customerDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE constraint failed")
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_CreateCustomer_NetworkTimeout(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		Email:     "joao.silva@email.com",
		CPF:       "123.456.789-00",
	}

	expectedError := errors.New("network timeout")
	mockRepo.On("CreateCustomer", customerDTO).Return(expectedError)

	service := NewCustomerService(mockRepo)

	err := service.CreateCustomer(customerDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "network timeout")
	mockRepo.AssertExpectations(t)
}

func TestCustomerService_GetCustomerByCpf_NetworkTimeout(t *testing.T) {
	mockRepo := &MockCustomerRepository{}

	cpf := "123.456.789-00"
	expectedError := errors.New("network timeout")

	mockRepo.On("GetCustomerByCpf", cpf).Return(nil, expectedError)

	service := NewCustomerService(mockRepo)

	customer, err := service.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Contains(t, err.Error(), "network timeout")
	mockRepo.AssertExpectations(t)
}
