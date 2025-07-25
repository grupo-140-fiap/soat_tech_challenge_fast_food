package gateways

import (
	"database/sql"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output"
)

type customerGateway struct {
	db *sql.DB
}

func NewCustomerGateway(db *sql.DB) output.CustomerGateway {
	return &customerGateway{
		db: db,
	}
}

func (g *customerGateway) Create(customer *entities.Customer) error {
	query := `
		INSERT INTO customers (first_name, last_name, cpf, email, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := g.db.Exec(query,
		customer.FirstName,
		customer.LastName,
		customer.CPF,
		customer.Email,
		customer.CreatedAt,
		customer.UpdatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	customer.ID = uint64(id)
	return nil
}

func (g *customerGateway) GetByCPF(cpf string) (*entities.Customer, error) {
	query := `
		SELECT id, first_name, last_name, cpf, email, created_at, updated_at
		FROM customers
		WHERE cpf = ?
	`

	row := g.db.QueryRow(query, cpf)

	var customer entities.Customer
	var createdAt, updatedAt string

	err := row.Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.CPF,
		&customer.Email,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	customer.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	customer.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &customer, nil
}

func (g *customerGateway) GetByID(id uint64) (*entities.Customer, error) {
	query := `
		SELECT id, first_name, last_name, cpf, email, created_at, updated_at
		FROM customers
		WHERE id = ?
	`

	row := g.db.QueryRow(query, id)

	var customer entities.Customer
	var createdAt, updatedAt string

	err := row.Scan(
		&customer.ID,
		&customer.FirstName,
		&customer.LastName,
		&customer.CPF,
		&customer.Email,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	customer.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	customer.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &customer, nil
}

func (g *customerGateway) Update(customer *entities.Customer) error {
	query := `
		UPDATE customers
		SET first_name = ?, last_name = ?, email = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := g.db.Exec(query,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.UpdatedAt,
		customer.ID,
	)

	return err
}

func (g *customerGateway) Delete(id uint64) error {
	query := `DELETE FROM customers WHERE id = ?`
	_, err := g.db.Exec(query, id)
	return err
}
