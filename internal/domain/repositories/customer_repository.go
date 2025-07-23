package repositories

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"

// CustomerRepository defines the contract for customer data persistence
// Following the Dependency Inversion Principle, this interface is defined
// in the domain layer but implemented in the infrastructure layer
type CustomerRepository interface {
	Create(customer *entities.Customer) error
	GetByCPF(cpf string) (*entities.Customer, error)
	GetByID(id uint64) (*entities.Customer, error)
	Update(customer *entities.Customer) error
	Delete(id uint64) error
}
