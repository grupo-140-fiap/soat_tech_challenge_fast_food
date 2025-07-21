package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

// CreatePayment godoc
// @Summary      Create a new payment
// @Description  Create a new payment using MercadoPago integration
// @Tags         payment
// @Accept       json
// @Produce      json
// @Param        payment  body      dto.PaymentDTO  true  "Payment data"
// @Success      200      {object}  object{message=string,Qrcode=string}
// @Router       /payments [post]
func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var paymentDTO dto.PaymentDTO

	if err := c.ShouldBindJSON(&paymentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})

		return
	}

	err := h.paymentService.CreatePayment(&paymentDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create payment",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment created successfully",
		"Qrcode":  paymentDTO.QrcodeUrl,
	})
}
