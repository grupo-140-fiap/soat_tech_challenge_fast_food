package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
)

type PaymentService interface {
	CreatePayment(payment *dto.PaymentDTO) error
}
