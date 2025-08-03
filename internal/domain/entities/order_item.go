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

func NewOrderItem(orderID, productID uint64, quantity uint32, price float32) *OrderItem {
	return &OrderItem{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (oi *OrderItem) UpdateQuantity(quantity uint32) {
	oi.Quantity = quantity
	oi.UpdatedAt = time.Now()
}

func (oi *OrderItem) CalculateSubtotal() float32 {
	return oi.Price * float32(oi.Quantity)
}

func (oi *OrderItem) IsValid() bool {
	return oi.ProductID > 0 && oi.Quantity > 0 && oi.Price > 0
}
