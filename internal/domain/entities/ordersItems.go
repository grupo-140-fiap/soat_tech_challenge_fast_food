package entities

import "time"

type OrdersItems struct {
	ID        uint64    `json:"id"`
	OrderId   uint64    `json:"order_id"`
	ProductId uint64    `json:"product_id"`
	Quantity  uint32    `json:"quantity"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
