package entities

import "time"

// OrderItem represents an item in an order.
type OrderItem struct {
	ID        uint64    `json:"id"`
	OrderID   uint64    `json:"order_id"`
	ProductID uint64    `json:"product_id"`
	Quantity  uint32    `json:"quantity"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewOrderItem creates a new order item
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

// UpdateQuantity updates the quantity of the order item
func (oi *OrderItem) UpdateQuantity(quantity uint32) {
	oi.Quantity = quantity
	oi.UpdatedAt = time.Now()
}

// CalculateSubtotal calculates the subtotal for this item
func (oi *OrderItem) CalculateSubtotal() float32 {
	return oi.Price * float32(oi.Quantity)
}

// IsValid validates order item business rules
func (oi *OrderItem) IsValid() bool {
	return oi.ProductID > 0 && oi.Quantity > 0 && oi.Price > 0
}
