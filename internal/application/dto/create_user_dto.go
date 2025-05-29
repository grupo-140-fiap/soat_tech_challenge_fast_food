package dto

type CreateCustomerDTO struct {
	FirstName string `json:"first_name" example:"João"`
	LastName  string `json:"last_name" example:"Silva"`
	Email     string `json:"email" example:"joao.silva@email.com"`
	CPF       string `json:"cpf" example:"123.456.789-00"`
}
