package dto

// CreatePaymentRequest represents the request to create a payment
type CreatePaymentRequest struct {
	OrderID       uint64  `json:"order_id" binding:"required"`
	Amount        float32 `json:"amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
}

// PaymentResponse represents the payment response
type PaymentResponse struct {
	ID            uint64  `json:"id"`
	OrderID       uint64  `json:"order_id"`
	Amount        float32 `json:"amount"`
	Status        string  `json:"status"`
	PaymentMethod string  `json:"payment_method"`
	TransactionID string  `json:"transaction_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// PaymentStatusResponse represents the payment status query response
type PaymentStatusResponse struct {
	ID            uint64  `json:"id"`
	OrderID       uint64  `json:"order_id"`
	Status        string  `json:"status"`
	Amount        float32 `json:"amount"`
	TransactionID string  `json:"transaction_id"`
	UpdatedAt     string  `json:"updated_at"`
}

// WebhookPaymentRequest represents the webhook payload for payment status updates
type WebhookPaymentRequest struct {
	TransactionID string  `json:"transaction_id" binding:"required"`
	OrderID       uint64  `json:"order_id" binding:"required"`
	Status        string  `json:"status" binding:"required"`
	Amount        float32 `json:"amount" binding:"required,gt=0"`
	PaymentMethod string  `json:"payment_method"`
	Timestamp     string  `json:"timestamp"`
}
