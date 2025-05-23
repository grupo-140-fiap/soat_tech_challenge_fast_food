package dto

type ProductDTO struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Category    string `json:"category"`
	Image       string `json:"image"`
}

//curl -X POST http://localhost:8080/api/v1/products -H "Content-Type: application/json" -d '{"name":"Pizza","description":"queijo","price":"40","category":"burger"}'
