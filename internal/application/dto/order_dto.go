package dto

type OrderDTO struct {
	ID         uint64         `json:"id"`
	CustomerId uint64         `json:"customer_id"`
	CPF        string         `json:"cpf"`
	Status     string         `json:"status"`
	Items      []OrderItemDTO `json:"items"`
}
