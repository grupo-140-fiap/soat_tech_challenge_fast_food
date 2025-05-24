package dto

type OrderDTO struct {
	ID         uint64         `json:"id"`
	CustomerId uint64         `json:"customer_id"`
	CPF        string         `json:"cpf"`
	Status     string         `json:"status"`
	Items      []OrderItemDTO `json:"items"`
}

// curl -X POST http://localhost:8080/api/v1/checkout -H "Content-Type: application/json" -d '{"customer_id":1,"cpf":"xxx.xxx.xxx","status":"received", "items":[{"order_id":1,"product_id":1,"quantity":1, "price": 5.66},{"order_id":1,"product_id":2,"quantity":1, "price": 2.88}]}'
