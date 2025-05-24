package entities

import "time"

type OrderItem struct {
	ID        uint64    `json:"id"`
	OrderID   uint64    `json:"order_id"`
	ProductID uint64    `json:"product_id"`
	Quantity  uint32    `json:"quantity"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
