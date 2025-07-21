package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
)

type CustomerService interface {
	CreateCustomer(customer *dto.CreateCustomerDTO) error
	GetCustomerByCpf(cpf string) (*entities.Customer, error)
}
