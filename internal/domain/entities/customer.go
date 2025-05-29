package entities

import "time"

// Customer represents a customer entity.
// swagger:model Customer
type Customer struct {
	ID        uint64    `json:"id" example:"1"`
	FirstName string    `json:"first_name" example:"João"`
	LastName  string    `json:"last_name" example:"Silva"`
	CPF       string    `json:"cpf" example:"123.456.789-00"`
	Email     string    `json:"email" example:"joao.silva@email.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-06-02T15:30:00Z"`
}
