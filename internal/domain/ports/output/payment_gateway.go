package output

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"

// PaymentGateway defines the contract for payment data access operations
type PaymentGateway interface {
	Create(payment *entities.Payment) error
	GetByID(id uint64) (*entities.Payment, error)
	GetByOrderID(orderID uint64) (*entities.Payment, error)
	GetByTransactionID(transactionID string) (*entities.Payment, error)
	Update(payment *entities.Payment) error
	Delete(id uint64) error
}
