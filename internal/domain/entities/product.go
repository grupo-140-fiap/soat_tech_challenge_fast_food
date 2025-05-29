package entities

import "time"

// Product represents a product entity.
// swagger:model Product
type Product struct {
	ID          uint64    `json:"id" example:"1"`
	Name        string    `json:"name" example:"Cheeseburger"`
	Description string    `json:"description" example:"Delicious cheeseburger with cheddar and pickles"`
	Price       float32   `json:"price" example:"12.99"`
	Category    string    `json:"category" example:"Sandwich"`
	Image       string    `json:"image" example:"https://example.com/images/cheeseburger.png"`
	CreatedAt   time.Time `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2024-06-01T12:00:00Z"`
}
