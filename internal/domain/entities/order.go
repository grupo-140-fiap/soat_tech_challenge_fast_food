package entities

import "time"

type Order struct {
	ID         uint64      `json:"id"`
	CustomerId uint64      `json:"customer_id"`
	CPF        string      `json:"cpf"`
	Status     OrderStatus `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Items      []OrderItem `json:"items"`
}

type OrderStatus string

const (
	OrderReceived   OrderStatus = "received"
	OrderInProgress OrderStatus = "in_progress"
	OrderReady      OrderStatus = "ready"
	OrderCompleted  OrderStatus = "completed"
	OrderCancelled  OrderStatus = "cancelled"
)

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

func (o *Order) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
	o.UpdatedAt = time.Now()
}

func (o *Order) UpdateStatus(status OrderStatus) {
	o.Status = status
	o.UpdatedAt = time.Now()
}

func (o *Order) CalculateTotal() float32 {
	var total float32
	for _, item := range o.Items {
		total += item.Price * float32(item.Quantity)
	}
	return total
}

func (o *Order) IsValid() bool {
	return len(o.Items) > 0
}
