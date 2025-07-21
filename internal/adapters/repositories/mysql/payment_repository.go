package persistance

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/order"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type PaymentRepository struct {
	payClient order.Client
}

func NewPaymentRepository(payClient order.Client) repositories.PaymentRepository {
	return &PaymentRepository{payClient: payClient}
}

func (u *PaymentRepository) CreatePayment(payment *dto.PaymentDTO) error {
	request := order.Request{
		Type:              "online",
		TotalAmount:       payment.Amount,
		ExternalReference: "ext_ref_1234",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: payment.Amount,
					PaymentMethod: &order.PaymentMethodRequest{
						ID:   "pix",
						Type: "bank_transfer",
					},
				},
			},
		},
		Payer: &order.PayerRequest{
			Email: payment.Email,
		},
	}

	resource, err := u.payClient.Create(context.Background(), request)
	if err != nil {
		return err
	}

	for _, paymentUrl := range resource.Transactions.Payments {
		payment.QrcodeUrl = paymentUrl.PaymentMethod.TicketURL
		return nil
	}

	return fmt.Errorf("Error payment mercado pago")
}
