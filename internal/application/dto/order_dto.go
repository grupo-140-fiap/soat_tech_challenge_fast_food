package dto

type OrderDTO struct {
	ID         uint64         `json:"id" example:"1"`
	CustomerId uint64         `json:"customer_id" example:"123"`
	CPF        string         `json:"cpf" example:"123.456.789-00"`
	Status     string         `json:"status" example:"pending"`
	Items      []OrderItemDTO `json:"items"`
}
