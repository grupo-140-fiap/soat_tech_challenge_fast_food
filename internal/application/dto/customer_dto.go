package dto

// CreateCustomerRequest represents the request to create a new customer
type CreateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required" example:"João"`
	LastName  string `json:"last_name" binding:"required" example:"Silva"`
	CPF       string `json:"cpf" binding:"required" example:"123.456.789-00"`
	Email     string `json:"email" binding:"required,email" example:"joao.silva@email.com"`
}

// CustomerResponse represents the response for customer operations
type CustomerResponse struct {
	ID        uint64 `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"João"`
	LastName  string `json:"last_name" example:"Silva"`
	CPF       string `json:"cpf" example:"123.456.789-00"`
	Email     string `json:"email" example:"joao.silva@email.com"`
	CreatedAt string `json:"created_at" example:"2024-06-01T12:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2024-06-02T15:30:00Z"`
}

// UpdateCustomerRequest represents the request to update a customer
type UpdateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required" example:"João"`
	LastName  string `json:"last_name" binding:"required" example:"Silva"`
	Email     string `json:"email" binding:"required,email" example:"joao.silva@email.com"`
}
