package gateways

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output"
)

type paymentGateway struct {
	db *sql.DB
}

func NewPaymentGateway(db *sql.DB) output.PaymentGateway {
	return &paymentGateway{
		db: db,
	}
}

func (g *paymentGateway) Create(payment *entities.Payment) error {
	query := `
		INSERT INTO payments (order_id, amount, status, payment_method, transaction_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := g.db.Exec(query,
		payment.OrderID,
		payment.Amount,
		string(payment.Status),
		payment.PaymentMethod,
		payment.TransactionID,
		payment.CreatedAt,
		payment.UpdatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	payment.ID = uint64(id)
	return nil
}

func (g *paymentGateway) GetByID(id uint64) (*entities.Payment, error) {
	query := `
		SELECT id, order_id, amount, status, payment_method, transaction_id, created_at, updated_at
		FROM payments
		WHERE id = ?
	`

	row := g.db.QueryRow(query, id)

	var payment entities.Payment
	var amount string
	var status string
	var createdAt, updatedAt string

	err := row.Scan(
		&payment.ID,
		&payment.OrderID,
		&amount,
		&status,
		&payment.PaymentMethod,
		&payment.TransactionID,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Parse amount
	if amountFloat, err := strconv.ParseFloat(amount, 32); err == nil {
		payment.Amount = float32(amountFloat)
	}

	// Parse status
	payment.Status = entities.PaymentStatus(status)

	// Parse timestamps
	payment.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	payment.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &payment, nil
}

func (g *paymentGateway) GetByOrderID(orderID uint64) (*entities.Payment, error) {
	query := `
		SELECT id, order_id, amount, status, payment_method, transaction_id, created_at, updated_at
		FROM payments
		WHERE order_id = ?
	`

	row := g.db.QueryRow(query, orderID)

	var payment entities.Payment
	var amount string
	var status string
	var createdAt, updatedAt string

	err := row.Scan(
		&payment.ID,
		&payment.OrderID,
		&amount,
		&status,
		&payment.PaymentMethod,
		&payment.TransactionID,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Parse amount
	if amountFloat, err := strconv.ParseFloat(amount, 32); err == nil {
		payment.Amount = float32(amountFloat)
	}

	// Parse status
	payment.Status = entities.PaymentStatus(status)

	// Parse timestamps
	payment.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	payment.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &payment, nil
}

func (g *paymentGateway) GetByTransactionID(transactionID string) (*entities.Payment, error) {
	query := `
		SELECT id, order_id, amount, status, payment_method, transaction_id, created_at, updated_at
		FROM payments
		WHERE transaction_id = ?
	`

	row := g.db.QueryRow(query, transactionID)

	var payment entities.Payment
	var amount string
	var status string
	var createdAt, updatedAt string

	err := row.Scan(
		&payment.ID,
		&payment.OrderID,
		&amount,
		&status,
		&payment.PaymentMethod,
		&payment.TransactionID,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Parse amount
	if amountFloat, err := strconv.ParseFloat(amount, 32); err == nil {
		payment.Amount = float32(amountFloat)
	}

	// Parse status
	payment.Status = entities.PaymentStatus(status)

	// Parse timestamps
	payment.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	payment.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &payment, nil
}

func (g *paymentGateway) Update(payment *entities.Payment) error {
	query := `
		UPDATE payments
		SET amount = ?, status = ?, payment_method = ?, transaction_id = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := g.db.Exec(query,
		payment.Amount,
		string(payment.Status),
		payment.PaymentMethod,
		payment.TransactionID,
		payment.UpdatedAt,
		payment.ID,
	)

	return err
}

func (g *paymentGateway) Delete(id uint64) error {
	query := `DELETE FROM payments WHERE id = ?`
	_, err := g.db.Exec(query, id)
	return err
}
