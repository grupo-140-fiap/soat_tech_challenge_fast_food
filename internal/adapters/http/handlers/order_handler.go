package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
)

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderDTO dto.OrderDTO

	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})

		return
	}

	err := h.orderService.CreateOrder(&orderDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
	})
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	id := c.Param("id")

	order, err := h.orderService.GetOrderById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find order",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, order)
}
