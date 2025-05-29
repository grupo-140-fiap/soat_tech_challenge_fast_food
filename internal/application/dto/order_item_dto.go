package dto

type OrderItemDTO struct {
	ID        uint64  `json:"id" example:"1"`
	OrderID   uint64  `json:"order_id" example:"100"`
	ProductId uint64  `json:"product_id" example:"200"`
	Quantity  uint64  `json:"quantity" example:"2"`
	Price     float32 `json:"price" example:"19.99"`
}
