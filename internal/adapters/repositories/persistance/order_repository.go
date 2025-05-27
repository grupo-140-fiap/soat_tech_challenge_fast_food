package persistance

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
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

	order.ID = uint64(lastID)

	fmt.Printf("Inserted record with ID: %d\n", lastID)

	return nil
}

func (u *OrderRepository) GetOrderById(id string) (*entities.Order, error) {
	query := "SELECT id, customer_id, cpf, status FROM orders WHERE id = ?"
	row := u.db.QueryRow(query, id)

	var orders entities.Order
	err := row.Scan(&orders.ID, &orders.CustomerId, &orders.CPF, &orders.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("orders with ID %s not found", id)
		}

		return nil, err
	}

	return &orders, nil
}

func (u *OrderRepository) UpdateOrderStatus(id string, status string) error {
	query := "UPDATE orders SET status = ? WHERE id = ?"

	_, err := u.db.Exec(query, status, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *OrderRepository) GetActiveOrders() (*[]entities.Order, error) {
	query := `
        SELECT o.id, o.customer_id, o.cpf, o.status, 
               oi.id, oi.product_id, oi.quantity, oi.price
        FROM orders o
        LEFT JOIN order_items oi ON o.id = oi.order_id
        WHERE o.status IN ('received', 'preparation', 'ready')
        ORDER BY 
            CASE 
                WHEN o.status = 'received' THEN 1
                WHEN o.status = 'preparation' THEN 2
                WHEN o.status = 'ready' THEN 3
            END,
            o.created_at ASC`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query active orders: %w", err)
	}
	defer rows.Close()

	orderMap := make(map[uint64]*entities.Order)

	for rows.Next() {
		var order entities.Order
		var customerID sql.NullInt64
		var itemID sql.NullInt64
		var productID sql.NullInt64
		var quantity sql.NullInt64
		var price sql.NullFloat64

		err := rows.Scan(
			&order.ID,
			&customerID,
			&order.CPF,
			&order.Status,
			&itemID,
			&productID,
			&quantity,
			&price,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning order row: %w", err)
		}

		if customerID.Valid {
			order.CustomerId = uint64(customerID.Int64)
		} else {
			order.CustomerId = 0
		}

		existingOrder, exists := orderMap[order.ID]
		if !exists {
			order.Items = []entities.OrderItem{}
			orderMap[order.ID] = &order
			existingOrder = &order
		}

		if itemID.Valid && productID.Valid {
			item := entities.OrderItem{
				ID:        uint64(itemID.Int64),
				OrderID:   order.ID,
				ProductID: uint64(productID.Int64),
				Quantity:  uint32(quantity.Int64),
				Price:     float32(price.Float64),
			}
			existingOrder.Items = append(existingOrder.Items, item)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating order rows: %w", err)
	}

	var orders []entities.Order
	for _, order := range orderMap {
		orders = append(orders, *order)
	}

	return &orders, nil
}
