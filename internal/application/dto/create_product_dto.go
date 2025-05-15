package dto

type CreateProductDTO struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Image       string `json:"image"`
}
