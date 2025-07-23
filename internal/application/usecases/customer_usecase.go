package usecases

import (
	"errors"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/repositories"
)

// CustomerUseCase defines the interface for customer business operations
// Following Interface Segregation Principle
type CustomerUseCase interface {
	CreateCustomer(request *dto.CreateCustomerRequest) (*dto.CustomerResponse, error)
	GetCustomerByCPF(cpf string) (*dto.CustomerResponse, error)
	GetCustomerByID(id uint64) (*dto.CustomerResponse, error)
	UpdateCustomer(id uint64, request *dto.UpdateCustomerRequest) (*dto.CustomerResponse, error)
	DeleteCustomer(id uint64) error
}

// customerUseCase implements CustomerUseCase interface
type customerUseCase struct {
	customerRepo repositories.CustomerRepository
}

// NewCustomerUseCase creates a new customer use case
func NewCustomerUseCase(customerRepo repositories.CustomerRepository) CustomerUseCase {
	return &customerUseCase{
		customerRepo: customerRepo,
	}
}

// CreateCustomer creates a new customer
func (uc *customerUseCase) CreateCustomer(request *dto.CreateCustomerRequest) (*dto.CustomerResponse, error) {
	// Business logic: Create domain entity
	customer := entities.NewCustomer(request.FirstName, request.LastName, request.CPF, request.Email)

	// Business validation
	if !customer.IsValid() {
		return nil, errors.New("invalid customer data")
	}

	// Check if customer already exists
	existingCustomer, _ := uc.customerRepo.GetByCPF(request.CPF)
	if existingCustomer != nil {
		return nil, errors.New("customer with this CPF already exists")
	}

	// Persist entity
	err := uc.customerRepo.Create(customer)
	if err != nil {
		return nil, err
	}

	// Return response DTO
	return &dto.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CPF:       customer.CPF,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: customer.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// GetCustomerByCPF retrieves a customer by CPF
func (uc *customerUseCase) GetCustomerByCPF(cpf string) (*dto.CustomerResponse, error) {
	customer, err := uc.customerRepo.GetByCPF(cpf)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return &dto.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CPF:       customer.CPF,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: customer.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// GetCustomerByID retrieves a customer by ID
func (uc *customerUseCase) GetCustomerByID(id uint64) (*dto.CustomerResponse, error) {
	customer, err := uc.customerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("customer not found")
	}

	return &dto.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CPF:       customer.CPF,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: customer.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// UpdateCustomer updates an existing customer
func (uc *customerUseCase) UpdateCustomer(id uint64, request *dto.UpdateCustomerRequest) (*dto.CustomerResponse, error) {
	customer, err := uc.customerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, errors.New("customer not found")
	}

	// Update entity using domain method
	customer.UpdateCustomer(request.FirstName, request.LastName, request.Email)

	// Business validation
	if !customer.IsValid() {
		return nil, errors.New("invalid customer data")
	}

	// Persist changes
	err = uc.customerRepo.Update(customer)
	if err != nil {
		return nil, err
	}

	return &dto.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		CPF:       customer.CPF,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt: customer.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// DeleteCustomer deletes a customer
func (uc *customerUseCase) DeleteCustomer(id uint64) error {
	customer, err := uc.customerRepo.GetByID(id)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	return uc.customerRepo.Delete(id)
}
