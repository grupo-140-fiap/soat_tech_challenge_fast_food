package input

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"

// PaymentUseCase defines the contract for payment business operations
type PaymentUseCase interface {
	CreatePayment(request *dto.CreatePaymentRequest) (*dto.PaymentResponse, error)
	GetPaymentStatus(orderID uint64) (*dto.PaymentStatusResponse, error)
	GetPaymentByTransactionID(transactionID string) (*dto.PaymentResponse, error)
	ProcessWebhookPayment(request *dto.WebhookPaymentRequest) error
}
