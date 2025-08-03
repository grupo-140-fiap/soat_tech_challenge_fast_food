package input

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"

// CustomerUseCase defines the contract for customer business operations
type CustomerUseCase interface {
	CreateCustomer(request *dto.CreateCustomerRequest) (*dto.CustomerResponse, error)
	GetCustomerByCPF(cpf string) (*dto.CustomerResponse, error)
	GetCustomerByID(id uint64) (*dto.CustomerResponse, error)
	UpdateCustomer(id uint64, request *dto.UpdateCustomerRequest) (*dto.CustomerResponse, error)
	DeleteCustomer(id uint64) error
}
