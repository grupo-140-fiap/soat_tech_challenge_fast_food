package repositories

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"

type PaymentRepository interface {
	CreatePayment(payment *dto.PaymentDTO) error
}
