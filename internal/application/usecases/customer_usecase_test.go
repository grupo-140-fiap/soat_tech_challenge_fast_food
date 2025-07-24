package usecases

import (
	"errors"
	"testing"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

type MockCustomerRepository struct {
	customers map[string]*entities.Customer
	nextID    uint64
}

func NewMockCustomerRepository() *MockCustomerRepository {
	return &MockCustomerRepository{
		customers: make(map[string]*entities.Customer),
		nextID:    1,
	}
}

func (m *MockCustomerRepository) Create(customer *entities.Customer) error {
	if _, exists := m.customers[customer.CPF]; exists {
		return errors.New("customer already exists")
	}
	customer.ID = m.nextID
	m.nextID++
	m.customers[customer.CPF] = customer
	return nil
}

func (m *MockCustomerRepository) GetByCPF(cpf string) (*entities.Customer, error) {
	customer, exists := m.customers[cpf]
	if !exists {
		return nil, nil
	}
	return customer, nil
}

func (m *MockCustomerRepository) GetByID(id uint64) (*entities.Customer, error) {
	for _, customer := range m.customers {
		if customer.ID == id {
			return customer, nil
		}
	}
	return nil, nil
}

func (m *MockCustomerRepository) Update(customer *entities.Customer) error {
	if _, exists := m.customers[customer.CPF]; !exists {
		return errors.New("customer not found")
	}
	m.customers[customer.CPF] = customer
	return nil
}

func (m *MockCustomerRepository) Delete(id uint64) error {
	for cpf, customer := range m.customers {
		if customer.ID == id {
			delete(m.customers, cpf)
			return nil
		}
	}
	return errors.New("customer not found")
}

func TestCustomerUseCase_CreateCustomer(t *testing.T) {
	// Arrange
	mockRepo := NewMockCustomerRepository()
	useCase := NewCustomerUseCase(mockRepo)

	request := &dto.CreateCustomerRequest{
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao@email.com",
	}

	// Act
	response, err := useCase.CreateCustomer(request)

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response == nil {
		t.Error("Expected response, got nil")
	}

	if response.FirstName != "João" {
		t.Errorf("Expected firstName 'João', got %s", response.FirstName)
	}

	if response.CPF != "123.456.789-00" {
		t.Errorf("Expected CPF '123.456.789-00', got %s", response.CPF)
	}
}

func TestCustomerUseCase_CreateCustomer_DuplicateCPF(t *testing.T) {
	// Arrange
	mockRepo := NewMockCustomerRepository()
	useCase := NewCustomerUseCase(mockRepo)

	// Primeiro cliente
	customer1 := &entities.Customer{
		ID:        1,
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.Create(customer1)

	// Segundo cliente com mesmo CPF
	request := &dto.CreateCustomerRequest{
		FirstName: "Maria",
		LastName:  "Santos",
		CPF:       "123.456.789-00", // CPF duplicado
		Email:     "maria@email.com",
	}

	// Act
	response, err := useCase.CreateCustomer(request)

	// Assert
	if err == nil {
		t.Error("Expected error for duplicate CPF, got nil")
	}

	if response != nil {
		t.Error("Expected nil response for duplicate CPF")
	}

	if err.Error() != "customer with this CPF already exists" {
		t.Errorf("Expected specific error message, got %s", err.Error())
	}
}

func TestCustomerUseCase_GetCustomerByCPF(t *testing.T) {
	// Arrange
	mockRepo := NewMockCustomerRepository()
	useCase := NewCustomerUseCase(mockRepo)

	// Criar cliente no mock
	customer := &entities.Customer{
		ID:        1,
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo.Create(customer)

	// Act
	response, err := useCase.GetCustomerByCPF("123.456.789-00")

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response == nil {
		t.Error("Expected response, got nil")
	}

	if response.FirstName != "João" {
		t.Errorf("Expected firstName 'João', got %s", response.FirstName)
	}
}

func TestCustomerUseCase_GetCustomerByCPF_NotFound(t *testing.T) {
	// Arrange
	mockRepo := NewMockCustomerRepository()
	useCase := NewCustomerUseCase(mockRepo)

	// Act
	response, err := useCase.GetCustomerByCPF("999.999.999-99")

	// Assert
	if err == nil {
		t.Error("Expected error for customer not found, got nil")
	}

	if response != nil {
		t.Error("Expected nil response for customer not found")
	}

	if err.Error() != "customer not found" {
		t.Errorf("Expected 'customer not found' error, got %s", err.Error())
	}
}
