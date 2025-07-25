package presenters

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type PaymentPresenter interface {
	PresentPayment(payment *dto.PaymentResponse) interface{}
	PresentPaymentStatus(status *dto.PaymentStatusResponse) interface{}
	PresentError(err error) interface{}
	PresentSuccess(message string) interface{}
}

type paymentPresenter struct{}

func NewPaymentPresenter() PaymentPresenter {
	return &paymentPresenter{}
}

func (p *paymentPresenter) PresentPayment(payment *dto.PaymentResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"data":    payment,
	}
}

func (p *paymentPresenter) PresentPaymentStatus(status *dto.PaymentStatusResponse) interface{} {
	return map[string]interface{}{
		"success": true,
		"data":    status,
	}
}

func (p *paymentPresenter) PresentError(err error) interface{} {
	return map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	}
}

func (p *paymentPresenter) PresentSuccess(message string) interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
	}
}
