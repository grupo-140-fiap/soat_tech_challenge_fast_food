package persistance

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/order"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
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
		TotalAmount:       "1.00",
		ExternalReference: "ext_ref_1234",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:   "pix",
						Type: "bank_transfer",
					},
				},
			},
		},
		Payer: &order.PayerRequest{
			Email: "test@testuser.com",
		},
	}

	resource, err := u.payClient.Create(context.Background(), request)
	if err != nil {
		return err
	}

	fmt.Println(resource)
	// resource.Transactions.Payments[0].PaymentMethod.TicketURL

	return nil
}
