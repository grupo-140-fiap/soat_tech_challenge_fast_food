package persistance

import (
    "database/sql"
	"fmt"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
    "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
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

func (u *CustomerRepository) GetCustomerByCpf(cpf string) (*entities.Customer, error) {
	query := "SELECT id, first_name, last_name, cpf, email FROM customers WHERE cpf = ?"
	row := u.db.QueryRow(query, cpf)

	var customer entities.Customer
	err := row.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.CPF, &customer.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer with CPF %s not found", cpf)
		}

		return nil, err
	}

	return &customer, nil
}