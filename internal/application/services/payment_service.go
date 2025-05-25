package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type PaymentService struct {
	paymentRepository repositories.PaymentRepository
}

func NewPaymentService(paymentRepository repositories.PaymentRepository) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepository,
	}
}

func (u *PaymentService) CreatePayment(payment *dto.PaymentDTO) error {
	err := u.paymentRepository.CreatePayment(payment)

	if err != nil {
		return err
	}

	return nil
}
