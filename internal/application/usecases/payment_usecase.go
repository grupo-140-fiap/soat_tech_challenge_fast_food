package usecases

import (
	"errors"
	"fmt"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output"
)

type paymentUseCase struct {
	paymentGateway output.PaymentGateway
	orderGateway   output.OrderGateway
}

func NewPaymentUseCase(paymentGateway output.PaymentGateway, orderGateway output.OrderGateway) input.PaymentUseCase {
	return &paymentUseCase{
		paymentGateway: paymentGateway,
		orderGateway:   orderGateway,
	}
}

func (uc *paymentUseCase) CreatePayment(request *dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {
	// Validate if order exists
	order, err := uc.orderGateway.GetByID(request.OrderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	if order == nil {
		return nil, errors.New("order not found")
	}

	// Check if payment already exists for this order
	existingPayment, _ := uc.paymentGateway.GetByOrderID(request.OrderID)
	if existingPayment != nil {
		// Return existing payment instead of creating a new one
		return &dto.PaymentResponse{
			ID:            existingPayment.ID,
			OrderID:       existingPayment.OrderID,
			Amount:        existingPayment.Amount,
			Status:        string(existingPayment.Status),
			PaymentMethod: existingPayment.PaymentMethod,
			TransactionID: existingPayment.TransactionID,
			CreatedAt:     existingPayment.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:     existingPayment.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}, nil
	}

	// Create payment only if it doesn't exist
	payment := entities.NewPayment(request.OrderID, request.Amount, request.PaymentMethod)

	err = uc.paymentGateway.Create(payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return &dto.PaymentResponse{
		ID:            payment.ID,
		OrderID:       payment.OrderID,
		Amount:        payment.Amount,
		Status:        string(payment.Status),
		PaymentMethod: payment.PaymentMethod,
		TransactionID: payment.TransactionID,
		CreatedAt:     payment.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     payment.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (uc *paymentUseCase) GetPaymentStatus(orderID uint64) (*dto.PaymentStatusResponse, error) {
	payment, err := uc.paymentGateway.GetByOrderID(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	if payment == nil {
		return nil, errors.New("payment not found for this order")
	}

	return &dto.PaymentStatusResponse{
		ID:            payment.ID,
		OrderID:       payment.OrderID,
		Status:        string(payment.Status),
		Amount:        payment.Amount,
		TransactionID: payment.TransactionID,
		UpdatedAt:     payment.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (uc *paymentUseCase) GetPaymentByTransactionID(transactionID string) (*dto.PaymentResponse, error) {
	payment, err := uc.paymentGateway.GetByTransactionID(transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	if payment == nil {
		return nil, errors.New("payment not found")
	}

	return &dto.PaymentResponse{
		ID:            payment.ID,
		OrderID:       payment.OrderID,
		Amount:        payment.Amount,
		Status:        string(payment.Status),
		PaymentMethod: payment.PaymentMethod,
		TransactionID: payment.TransactionID,
		CreatedAt:     payment.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:     payment.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

func (uc *paymentUseCase) ProcessWebhookPayment(request *dto.WebhookPaymentRequest) error {
	// Find payment by order ID
	payment, err := uc.paymentGateway.GetByOrderID(request.OrderID)
	if err != nil {
		return fmt.Errorf("failed to get payment: %w", err)
	}
	if payment == nil {
		return errors.New("payment not found for this order")
	}

	// Validate payment status
	status := entities.PaymentStatus(request.Status)
	if status != entities.PaymentStatusApproved &&
		status != entities.PaymentStatusRejected &&
		status != entities.PaymentStatusCanceled {
		return errors.New("invalid payment status received")
	}

	// Update payment status
	payment.UpdateStatus(status, request.TransactionID)

	err = uc.paymentGateway.Update(payment)
	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	// If payment is approved, update order status
	if payment.IsApproved() {
		order, err := uc.orderGateway.GetByID(payment.OrderID)
		if err != nil {
			return fmt.Errorf("failed to get order for status update: %w", err)
		}
		if order != nil {
			order.UpdateStatus(entities.OrderReceived)
			err = uc.orderGateway.Update(order)
			if err != nil {
				return fmt.Errorf("failed to update order status: %w", err)
			}
		}
	}

	return nil
}
