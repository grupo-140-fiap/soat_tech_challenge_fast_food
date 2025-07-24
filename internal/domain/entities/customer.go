package entities

import "time"

type Customer struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CPF       string    `json:"cpf"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCustomer(firstName, lastName, cpf, email string) *Customer {
	return &Customer{
		FirstName: firstName,
		LastName:  lastName,
		CPF:       cpf,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Customer) UpdateCustomer(firstName, lastName, email string) {
	c.FirstName = firstName
	c.LastName = lastName
	c.Email = email
	c.UpdatedAt = time.Now()
}

func (c *Customer) IsValid() bool {
	return c.FirstName != "" && c.LastName != "" && c.CPF != "" && c.Email != ""
}
