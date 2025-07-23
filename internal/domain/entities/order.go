package entities

import "time"

// Order represents an order entity in the domain layer.
type Order struct {
	ID         uint64      `json:"id"`
	CustomerId uint64      `json:"customer_id"`
	CPF        string      `json:"cpf"`
	Status     OrderStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Items      []OrderItem `json:"items"`
}

// OrderStatus represents valid order statuses
type OrderStatus string

const (
	OrderReceived   OrderStatus = "received"
	OrderInProgress OrderStatus = "in_progress"
	OrderReady      OrderStatus = "ready"
	OrderCompleted  OrderStatus = "completed"
	OrderCancelled  OrderStatus = "cancelled"
)

// NewOrder creates a new order
func NewOrder(customerId uint64, cpf string) *Order {
	return &Order{
		CustomerId: customerId,
		CPF:        cpf,
		Status:     OrderReceived,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Items:      make([]OrderItem, 0),
	}
}

// AddItem adds an item to the order
func (o *Order) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
	o.UpdatedAt = time.Now()
}

// UpdateStatus updates the order status
func (o *Order) UpdateStatus(status OrderStatus) {
	o.Status = status
	o.UpdatedAt = time.Now()
}

// CalculateTotal calculates the total order value
func (o *Order) CalculateTotal() float32 {
	var total float32
	for _, item := range o.Items {
		total += item.Price * float32(item.Quantity)
	}
	return total
}

// IsValid validates order business rules
func (o *Order) IsValid() bool {
	return len(o.Items) > 0 && o.CPF != ""
}
