package entities

import "time"

type Product struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Category    string    `json:"category"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductCategory string

const (
	SnackCategory   ProductCategory = "snack"
	DrinkCategory   ProductCategory = "drink"
	DessertCategory ProductCategory = "dessert"
	SideCategory    ProductCategory = "side"
)

func NewProduct(name, description string, price float32, category ProductCategory, imageUrl string) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Category:    string(category),
		ImageUrl:    imageUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (p *Product) UpdateProduct(name, description string, price float32, category ProductCategory, imageUrl string) {
	p.Name = name
	p.Description = description
	p.Price = price
	p.Category = string(category)
	p.ImageUrl = imageUrl
	p.UpdatedAt = time.Now()
}

func (p *Product) IsValid() bool {
	return p.Name != "" && p.Price > 0 && p.Category != ""
}

func IsValidCategory(category string) bool {
	validCategories := []string{
		string(SnackCategory),
		string(DrinkCategory),
		string(DessertCategory),
		string(SideCategory),
	}

	for _, valid := range validCategories {
		if category == valid {
			return true
		}
	}
	return false
}
