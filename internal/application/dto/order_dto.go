package dto

type CreateOrderRequest struct {
	CustomerId uint64             `json:"customer_id" example:"123"`
	CPF        string             `json:"cpf" example:"123.456.789-00"`
	Items      []OrderItemRequest `json:"items" binding:"required,dive"`
}

type OrderItemRequest struct {
	ProductID uint64 `json:"product_id" binding:"required" example:"200"`
	Quantity  uint32 `json:"quantity" binding:"required,gt=0" example:"2"`
}

type OrderResponse struct {
	ID         uint64              `json:"id" example:"1"`
	CustomerId uint64              `json:"customer_id" example:"123"`
	CPF        string              `json:"cpf" example:"123.456.789-00"`
	Status     string              `json:"status" example:"received"`
	Items      []OrderItemResponse `json:"items"`
	Total      float32             `json:"total" example:"39.98"`
	CreatedAt  string              `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt  string              `json:"updated_at" example:"2024-06-01T12:30:00Z"`
}

type OrderItemResponse struct {
	ID        uint64  `json:"id" example:"1"`
	ProductID uint64  `json:"product_id" example:"200"`
	Quantity  uint32  `json:"quantity" example:"2"`
	Price     float32 `json:"price" example:"19.99"`
	Subtotal  float32 `json:"subtotal" example:"39.98"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required" example:"received" enums:"awaiting_payment,received,in_progress,ready,completed,cancelled"`
}
