package entities

import "time"

type Orders struct {
	ID         uint64    `json:"id"`
	CustomerId uint64    `json:"customer_id"`
	CPF        string    `json:"cpf"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
