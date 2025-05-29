package dto

type ProductDTO struct {
	ID          uint64 `json:"id" example:"1"`
	Name        string `json:"name" example:"Cheeseburger"`
	Description string `json:"description" example:"Delicioso cheeseburger com queijo cheddar e molho especial"`
	Price       string `json:"price" example:"19.90"`
	Category    string `json:"category" example:"Lanche"`
	Image       string `json:"image" example:"https://example.com/images/cheeseburger.png"`
}
