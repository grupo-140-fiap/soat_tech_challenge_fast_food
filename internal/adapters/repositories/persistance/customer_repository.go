package persistance

import (
    "database/sql"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type CustomerRepository struct {
    db *sql.DB
}

func NewCustomerRepository(db *sql.DB) repositories.CustomerRepository {
    return &CustomerRepository{db: db}
}

func (u *CustomerRepository) CreateCustomer(customer *dto.CreateCustomerDTO) error {
    query := "INSERT INTO customers (first_name, last_name, cpf, email) VALUES (?, ?, ?, ?)"

    _, err := u.db.Exec(query, customer.FirstName, customer.LastName, customer.CPF, customer.Email)

    if err != nil {
        return err
    }

    return nil
}