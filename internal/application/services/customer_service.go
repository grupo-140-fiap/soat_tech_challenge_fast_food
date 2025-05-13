package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
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