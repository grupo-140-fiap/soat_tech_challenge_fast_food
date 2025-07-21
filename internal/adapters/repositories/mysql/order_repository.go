package persistance

import (
	"database/sql"
	"fmt"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) repositories.OrderRepository {
	return &OrderRepository{db: db}
}

func (u *OrderRepository) GetOrders() ([]entities.Order, error) {
	rows, err := u.db.Query("SELECT id, customer_id, cpf, status, created_at, updated_at FROM orders ")
	if err != nil {
		return nil, fmt.Errorf("orders not found")
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var o entities.Order
		if err := rows.Scan(&o.ID, &o.CustomerId, &o.CPF, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}

		items, err := u.getOrderItems(o.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to load items for order %d: %w", o.ID, err)
		}
		o.Items = items

		orders = append(orders, o)
	}

	return orders, nil
}

func (u *OrderRepository) CreateOrder(order *entities.Order) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	result, err := tx.Exec(
		"INSERT INTO orders (customer_id, cpf, status) VALUES (?, ?, ?)",
		order.CustomerId, order.CPF, "received",
	)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get order ID: %w", err)
	}
	if len(order.Items) > 0 {
		stmt, err := tx.Prepare("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)")
		if err != nil {
			return fmt.Errorf("failed to prepare item statement: %w", err)
		}
		defer stmt.Close()

		for _, item := range order.Items {
			var productPrice float32
			err = tx.QueryRow("SELECT price FROM products WHERE id = ?", item.ProductId).Scan(&productPrice)
			if err != nil {
				return fmt.Errorf("failed to get price for product ID %d: %w", item.ProductId, err)
			}

			_, err = stmt.Exec(orderID, item.ProductId, item.Quantity, productPrice)
			if err != nil {
				return fmt.Errorf("failed to create order item (product_id: %d): %w", item.ProductId, err)
			}
		}
	}

	return nil
}

func (u *OrderRepository) GetOrderById(id string) (*entities.Order, error) {
	query := "SELECT id, customer_id, cpf, status, created_at, updated_at FROM orders WHERE id = ?"
	row := u.db.QueryRow(query, id)

	var orders entities.Order
	err := row.Scan(&orders.ID, &orders.CustomerId, &orders.CPF, &orders.Status, &orders.CreatedAt, &orders.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("orders with ID %s not found", id)
		}

		return nil, err
	}

	items, err := u.getOrderItems(orders.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to load items for order %d: %w", orders.ID, err)
	}
	orders.Items = items

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

func (u *OrderRepository) getOrderItems(orderID uint64) ([]entities.OrderItem, error) {
	query := "SELECT id, order_id, product_id, quantity, price, created_at, updated_at FROM order_items WHERE order_id = ?"
	rows, err := u.db.Query(query, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to query order items: %w", err)
	}
	defer rows.Close()

	var items []entities.OrderItem
	for rows.Next() {
		var item entities.OrderItem
		err := rows.Scan(&item.ID, &item.OrderID, &item.ProductID, &item.Quantity, &item.Price, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order item: %w", err)
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating order item rows: %w", err)
	}

	return items, nil
}
