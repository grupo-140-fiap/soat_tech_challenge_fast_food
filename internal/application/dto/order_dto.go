package dto

type OrderDTO struct {
	CustomerId uint64         `json:"customer_id" example:"123"`
	CPF        string         `json:"cpf" example:"123.456.789-00"`
	Status     string         `json:"status" example:"pending"`
	Items      []OrderItemDTO `json:"items"`
}
