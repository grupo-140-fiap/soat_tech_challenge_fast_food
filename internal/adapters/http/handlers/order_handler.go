package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
)

type OrderHandler struct {
	orderService services.OrderService
}

type RequestBody struct {
	Status string `json:"status"`
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// GetOrders godoc
// @Summary      Get all orders
// @Description  Retrieves a list of all orders
// @Tags         orders
// @Produce      json
// @Success      200  {array}   entities.Order
// @Router       /orders [get]
func (h *OrderHandler) GetOrders(c *gin.Context) {
	orders, err := h.orderService.GetOrders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find orders",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, orders)
}

// CreateOrder godoc
// @Summary      Create a new order
// @Description  Creates a new order using the provided JSON payload
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order  body      dto.OrderDTO  true  "Order data"
// @Success      200    {object}  map[string]interface{}  "Order created successfully"
// @Failure      400    {object}  map[string]interface{}  "Invalid input"
// @Failure      500    {object}  map[string]interface{}  "Failed to create order"
// @Router       /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var orderDTO dto.OrderDTO

	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})

		return
	}

	if len(orderDTO.Items) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Invalid input, no items",
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

// GetOrderById godoc
// @Summary      Get order by ID
// @Description  Retrieves an order by its unique identifier
// @Tags         orders
// @Produce      json
// @Param        id   path      string  true  "Order ID"
// @Success      200  {object}  entities.Order
// @Failure      500  {object}  map[string]interface{}  "Failed on find order"
// @Router       /orders/{id} [get]
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

// UpdateOrderStatus godoc
// @Summary      Update order status
// @Description  Updates the status of an existing order
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id      path    string      true  "Order ID"
// @Param        status  body    RequestBody  true  "New status"
// @Success      200     {object}  map[string]interface{}  "Order status updated successfully"
// @Failure      400     {object}  map[string]interface{}  "Invalid input"
// @Failure      500     {object}  map[string]interface{}  "Failed on update order status"
// @Router       /orders/{id}/status [patch]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	var body RequestBody
	var id = c.Param("id")
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})

		return
	}

	err := h.orderService.UpdateOrderStatus(id, body.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on update order status",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order status updated successfully",
	})
}
