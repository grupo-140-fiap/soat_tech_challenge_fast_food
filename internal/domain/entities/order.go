package entities

import "time"

// Order represents an order entity.
// swagger:model Order
type Order struct {
	ID         uint64      `json:"id" example:"1"`
	CustomerId uint64      `json:"customer_id" example:"123"`
	CPF        string      `json:"cpf" example:"123.456.789-00"`
	Status     string      `json:"status" example:"received"`
	CreatedAt  time.Time   `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt  time.Time   `json:"updated_at" example:"2024-06-01T12:30:00Z"`
	Items      []OrderItem `json:"items"`
}
