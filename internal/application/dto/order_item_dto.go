package dto

type OrderItemDTO struct {
	ID        uint64  `json:"id"`
	OrderId   uint64  `json:"order_id"`
	ProductId uint64  `json:"product_id"`
	Quantity  uint64  `json:"quantity"`
	Price     float32 `json:"price"`
}
