package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"
)

type PaymentController struct {
	paymentUseCase input.PaymentUseCase
	presenter      presenters.PaymentPresenter
}

func NewPaymentController(
	paymentUseCase input.PaymentUseCase,
	presenter presenters.PaymentPresenter,
) *PaymentController {
	return &PaymentController{
		paymentUseCase: paymentUseCase,
		presenter:      presenter,
	}
}

// CreatePayment godoc
// @Summary Create new payment
// @Description Create a new payment for an order
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body dto.CreatePaymentRequest true "payment"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments [post]
func (ctrl *PaymentController) CreatePayment(c *gin.Context) {
	var request dto.CreatePaymentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ctrl.presenter.PresentError(err))
		return
	}

	response, err := ctrl.paymentUseCase.CreatePayment(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ctrl.presenter.PresentError(err))
		return
	}

	c.JSON(http.StatusCreated, ctrl.presenter.PresentPayment(response))
}

// GetPaymentStatus godoc
// @Summary Get payment status by order ID
// @Description Get the current payment status for an order
// @Tags payments
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments/status/{order_id} [get]
func (ctrl *PaymentController) GetPaymentStatus(c *gin.Context) {
	orderIDStr := c.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, ctrl.presenter.PresentError(err))
		return
	}

	response, err := ctrl.paymentUseCase.GetPaymentStatus(orderID)
	if err != nil {
		if err.Error() == "payment not found for this order" {
			c.JSON(http.StatusNotFound, ctrl.presenter.PresentError(err))
			return
		}
		c.JSON(http.StatusInternalServerError, ctrl.presenter.PresentError(err))
		return
	}

	c.JSON(http.StatusOK, ctrl.presenter.PresentPaymentStatus(response))
}

// GetPaymentByTransactionID godoc
// @Summary Get payment by transaction ID
// @Description Get payment details by transaction ID
// @Tags payments
// @Produce json
// @Param transaction_id path string true "Transaction ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments/transaction/{transaction_id} [get]
func (ctrl *PaymentController) GetPaymentByTransactionID(c *gin.Context) {
	transactionID := c.Param("transaction_id")
	if transactionID == "" {
		c.JSON(http.StatusBadRequest, ctrl.presenter.PresentError(gin.Error{Err: gin.Error{Type: gin.ErrorTypeBind}.Err, Type: gin.ErrorTypeBind}))
		return
	}

	response, err := ctrl.paymentUseCase.GetPaymentByTransactionID(transactionID)
	if err != nil {
		if err.Error() == "payment not found" {
			c.JSON(http.StatusNotFound, ctrl.presenter.PresentError(err))
			return
		}
		c.JSON(http.StatusInternalServerError, ctrl.presenter.PresentError(err))
		return
	}

	c.JSON(http.StatusOK, ctrl.presenter.PresentPayment(response))
}

// PaymentWebhook godoc
// @Summary Payment webhook endpoint
// @Description Webhook endpoint to receive payment status updates from payment provider
// @Tags payments
// @Accept json
// @Produce json
// @Param webhook body dto.WebhookPaymentRequest true "webhook payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /payments/webhook [post]
func (ctrl *PaymentController) PaymentWebhook(c *gin.Context) {
	var request dto.WebhookPaymentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ctrl.presenter.PresentError(err))
		return
	}

	err := ctrl.paymentUseCase.ProcessWebhookPayment(&request)
	if err != nil {
		if err.Error() == "payment not found for this order" || err.Error() == "invalid payment status received" {
			c.JSON(http.StatusBadRequest, ctrl.presenter.PresentError(err))
			return
		}
		c.JSON(http.StatusInternalServerError, ctrl.presenter.PresentError(err))
		return
	}

	c.JSON(http.StatusOK, ctrl.presenter.PresentSuccess("Payment webhook processed successfully"))
}
