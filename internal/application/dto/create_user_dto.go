package dto

type CreateCustomerDTO struct {
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    CPF       string `json:"cpf"`
}