package dto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"Cheeseburger"`
	Description string  `json:"description" binding:"required" example:"Delicious cheeseburger with cheddar and pickles"`
	Price       float32 `json:"price" binding:"required,gt=0" example:"12.99"`
	Category    string  `json:"category" binding:"required" example:"snack"`
	ImageUrl    string  `json:"image_url" example:"https://example.com/images/cheeseburger.png"`
}

type ProductResponse struct {
	ID          uint64  `json:"id" example:"1"`
	Name        string  `json:"name" example:"Cheeseburger"`
	Description string  `json:"description" example:"Delicious cheeseburger with cheddar and pickles"`
	Price       float32 `json:"price" example:"12.99"`
	Category    string  `json:"category" example:"snack"`
	ImageUrl    string  `json:"image_url" example:"https://example.com/images/cheeseburger.png"`
	CreatedAt   string  `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt   string  `json:"updated_at" example:"2024-06-01T12:00:00Z"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required" example:"Cheeseburger"`
	Description string  `json:"description" binding:"required" example:"Delicious cheeseburger with cheddar and pickles"`
	Price       float32 `json:"price" binding:"required,gt=0" example:"12.99"`
	Category    string  `json:"category" binding:"required" example:"snack"`
	ImageUrl    string  `json:"image_url" example:"https://example.com/images/cheeseburger.png"`
}
