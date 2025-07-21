package dto

type PaymentDTO struct {
	Amount    string `json:"amount"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	QrcodeUrl string `json:"QrcodeUrl"`
}
