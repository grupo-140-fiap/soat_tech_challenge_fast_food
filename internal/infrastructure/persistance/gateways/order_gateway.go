package gateways

import (
	"database/sql"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/repositories"
)

type orderGateway struct {
	db *sql.DB
}

func NewOrderGateway(db *sql.DB) repositories.OrderRepository {
	return &orderGateway{
		db: db,
	}
}

func (g *orderGateway) Create(order *entities.Order) error {
	query := `
		INSERT INTO orders (customer_id, cpf, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := g.db.Exec(query,
		order.CustomerId,
		order.CPF,
		string(order.Status),
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	order.ID = uint64(id)
	return nil
}

func (g *orderGateway) GetByID(id uint64) (*entities.Order, error) {
	query := `
		SELECT id, customer_id, cpf, status, created_at, updated_at
		FROM orders
		WHERE id = ?
	`

	row := g.db.QueryRow(query, id)

	var order entities.Order
	var status string
	var createdAt, updatedAt string

	err := row.Scan(
		&order.ID,
		&order.CustomerId,
		&order.CPF,
		&status,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	order.Status = entities.OrderStatus(status)

	order.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	order.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &order, nil
}

func (g *orderGateway) GetByCPF(cpf string) ([]*entities.Order, error) {
	query := `
		SELECT id, customer_id, cpf, status, created_at, updated_at
		FROM orders
		WHERE cpf = ?
		ORDER BY created_at DESC
	`

	rows, err := g.db.Query(query, cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entities.Order

	for rows.Next() {
		var order entities.Order
		var status string
		var createdAt, updatedAt string

		err := rows.Scan(
			&order.ID,
			&order.CustomerId,
			&order.CPF,
			&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			continue
		}

		order.Status = entities.OrderStatus(status)

		// Parse timestamps
		order.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		order.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		orders = append(orders, &order)
	}

	return orders, nil
}

func (g *orderGateway) GetByCustomerID(customerID uint64) ([]*entities.Order, error) {
	query := `
		SELECT id, customer_id, cpf, status, created_at, updated_at
		FROM orders
		WHERE customer_id = ?
		ORDER BY created_at DESC
	`

	rows, err := g.db.Query(query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entities.Order

	for rows.Next() {
		var order entities.Order
		var status string
		var createdAt, updatedAt string

		err := rows.Scan(
			&order.ID,
			&order.CustomerId,
			&order.CPF,
			&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			continue
		}

		order.Status = entities.OrderStatus(status)

		// Parse timestamps
		order.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		order.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		orders = append(orders, &order)
	}

	return orders, nil
}

func (g *orderGateway) GetAll() ([]*entities.Order, error) {
	query := `
		SELECT id, customer_id, cpf, status, created_at, updated_at
		FROM orders
		ORDER BY created_at DESC
	`

	rows, err := g.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*entities.Order

	for rows.Next() {
		var order entities.Order
		var status string
		var createdAt, updatedAt string

		err := rows.Scan(
			&order.ID,
			&order.CustomerId,
			&order.CPF,
			&status,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			continue
		}

		order.Status = entities.OrderStatus(status)

		// Parse timestamps
		order.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		order.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		orders = append(orders, &order)
	}

	return orders, nil
}

func (g *orderGateway) Update(order *entities.Order) error {
	query := `
		UPDATE orders
		SET status = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := g.db.Exec(query,
		string(order.Status),
		order.UpdatedAt,
		order.ID,
	)

	return err
}

func (g *orderGateway) Delete(id uint64) error {
	_, err := g.db.Exec("DELETE FROM order_items WHERE order_id = ?", id)
	if err != nil {
		return err
	}

	query := `DELETE FROM orders WHERE id = ?`
	_, err = g.db.Exec(query, id)
	return err
}

type orderItemGateway struct {
	db *sql.DB
}

func NewOrderItemGateway(db *sql.DB) repositories.OrderItemRepository {
	return &orderItemGateway{
		db: db,
	}
}

func (g *orderItemGateway) Create(orderItem *entities.OrderItem) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity, price, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := g.db.Exec(query,
		orderItem.OrderID,
		orderItem.ProductID,
		orderItem.Quantity,
		orderItem.Price,
		orderItem.CreatedAt,
		orderItem.UpdatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	orderItem.ID = uint64(id)
	return nil
}

func (g *orderItemGateway) GetByOrderID(orderID uint64) ([]*entities.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity, price, created_at, updated_at
		FROM order_items
		WHERE order_id = ?
		ORDER BY id
	`

	rows, err := g.db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*entities.OrderItem

	for rows.Next() {
		var item entities.OrderItem
		var createdAt, updatedAt string

		err := rows.Scan(
			&item.ID,
			&item.OrderID,
			&item.ProductID,
			&item.Quantity,
			&item.Price,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			continue
		}

		item.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		item.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		items = append(items, &item)
	}

	return items, nil
}

func (g *orderItemGateway) Update(orderItem *entities.OrderItem) error {
	query := `
		UPDATE order_items
		SET quantity = ?, price = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := g.db.Exec(query,
		orderItem.Quantity,
		orderItem.Price,
		orderItem.UpdatedAt,
		orderItem.ID,
	)

	return err
}

func (g *orderItemGateway) Delete(id uint64) error {
	query := `DELETE FROM order_items WHERE id = ?`
	_, err := g.db.Exec(query, id)
	return err
}
