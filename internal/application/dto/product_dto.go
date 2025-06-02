package dto

type ProductDTO struct {
	Name        string `json:"name" example:"Cheeseburger"`
	Description string `json:"description" example:"Delicioso cheeseburger com queijo cheddar e molho especial"`
	Price       string `json:"price" example:"19.90"`
	Category    string `json:"category" example:"snack"`
	ImageUrl    string `json:"image_url" example:"https://example.com/images/cheeseburger.png"`
}
