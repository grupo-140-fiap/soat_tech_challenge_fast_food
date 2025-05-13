package entities

import "time"

type Customer struct {
    Id        uint64    `json:"id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    CPF       string    `json:"cpf"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}