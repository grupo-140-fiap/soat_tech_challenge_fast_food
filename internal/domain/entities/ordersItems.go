package entities

import "time"

// OrderItem represents an item in an order.
// swagger:model OrderItem
type OrderItem struct {
	ID        uint64    `json:"id" example:"1"`
	OrderID   uint64    `json:"order_id" example:"100"`
	ProductID uint64    `json:"product_id" example:"200"`
	Quantity  uint32    `json:"quantity" example:"2"`
	Price     float32   `json:"price" example:"19.99"`
	CreatedAt time.Time `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-06-01T12:30:00Z"`
}
