package entities

import (
	"time"
)

// PaymentStatus represents the possible payment statuses
type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "pending"
	PaymentStatusApproved PaymentStatus = "approved"
	PaymentStatusRejected PaymentStatus = "rejected"
	PaymentStatusCanceled PaymentStatus = "canceled"
)

// Payment represents a payment entity
type Payment struct {
	ID            uint64        `json:"id"`
	OrderID       uint64        `json:"order_id"`
	Amount        float32       `json:"amount"`
	Status        PaymentStatus `json:"status"`
	PaymentMethod string        `json:"payment_method"`
	TransactionID string        `json:"transaction_id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

// NewPayment creates a new payment instance
func NewPayment(orderID uint64, amount float32, paymentMethod string) *Payment {
	now := time.Now()
	return &Payment{
		OrderID:       orderID,
		Amount:        amount,
		Status:        PaymentStatusPending,
		PaymentMethod: paymentMethod,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// UpdateStatus updates the payment status and transaction ID
func (p *Payment) UpdateStatus(status PaymentStatus, transactionID string) {
	p.Status = status
	p.TransactionID = transactionID
	p.UpdatedAt = time.Now()
}

// IsApproved returns true if the payment is approved
func (p *Payment) IsApproved() bool {
	return p.Status == PaymentStatusApproved
}

// IsRejected returns true if the payment is rejected
func (p *Payment) IsRejected() bool {
	return p.Status == PaymentStatusRejected
}

// IsPending returns true if the payment is pending
func (p *Payment) IsPending() bool {
	return p.Status == PaymentStatusPending
}
