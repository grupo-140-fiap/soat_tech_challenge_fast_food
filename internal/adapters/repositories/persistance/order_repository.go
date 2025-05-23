package persistance

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) repositories.OrderRepository {
	return &OrderRepository{db: db}
}

func (u *OrderRepository) CreateOrder(order *dto.OrderDTO) error {
	// Prepare the INSERT statement
	stmt, err := u.db.Prepare("INSERT INTO orders (customer_id, cpf, status) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the statement and get the result
	result, err := stmt.Exec(order.CustomerId, order.CPF, order.Status)
	if err != nil {
		log.Fatal(err)
	}

	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted record with ID: %d\n", lastID)

	return nil
}
