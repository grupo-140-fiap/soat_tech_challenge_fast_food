package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type CustomerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (u *CustomerService) CreateCustomer(customer *dto.CreateCustomerDTO) error {
	err := u.customerRepository.CreateCustomer(customer)

	if err != nil {
		return err
	}

	return nil
}

func (u *CustomerService) GetCustomerByCpf(cpf string) (*entities.Customer, error) {
	customer, err := u.customerRepository.GetCustomerByCpf(cpf)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
