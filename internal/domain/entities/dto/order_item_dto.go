package dto

type OrderItemDTO struct {
	ProductId uint64 `json:"product_id" example:"1"`
	Quantity  uint64 `json:"quantity" example:"2"`
}
