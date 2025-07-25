package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/usecases"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"
)

type OrderController struct {
	orderUseCase usecases.OrderUseCase
	presenter    presenters.OrderPresenter
}

func NewOrderController(
	orderUseCase usecases.OrderUseCase,
	presenter presenters.OrderPresenter,
) *OrderController {
	return &OrderController{
		orderUseCase: orderUseCase,
		presenter:    presenter,
	}
}

// CreateOrder godoc
// @Summary Create new order
// @Description Create new order with items
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequest true "order"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders [post]
func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var request dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := ctrl.orderUseCase.CreateOrder(&request)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentOrder(order)
	c.JSON(http.StatusCreated, response)
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Get order by ID with items
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/{id} [get]
func (ctrl *OrderController) GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := ctrl.orderUseCase.GetOrderByID(id)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := ctrl.presenter.PresentOrder(order)
	c.JSON(http.StatusOK, response)
}

// GetOrdersByCPF godoc
// @Summary Get orders by CPF
// @Description Get all orders for a specific CPF
// @Tags orders
// @Produce json
// @Param cpf path string true "Customer CPF"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/cpf/{cpf} [get]
func (ctrl *OrderController) GetOrdersByCPF(c *gin.Context) {
	cpf := c.Param("cpf")

	orders, err := ctrl.orderUseCase.GetOrdersByCPF(cpf)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentOrders(orders)
	c.JSON(http.StatusOK, response)
}

// GetOrdersByCustomerID godoc
// @Summary Get orders by customer ID
// @Description Get all orders for a specific customer ID
// @Tags orders
// @Produce json
// @Param customerId path int true "Customer ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/customer/{customerId} [get]
func (ctrl *OrderController) GetOrdersByCustomerID(c *gin.Context) {
	customerIdStr := c.Param("customerId")
	customerId, err := strconv.ParseUint(customerIdStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	orders, err := ctrl.orderUseCase.GetOrdersByCustomerID(customerId)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentOrders(orders)
	c.JSON(http.StatusOK, response)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Get all orders in the system
// @Tags orders
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders [get]
func (ctrl *OrderController) GetAllOrders(c *gin.Context) {
	orders, err := ctrl.orderUseCase.GetAllOrders()
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentOrders(orders)
	c.JSON(http.StatusOK, response)
}

// GetOrdersForKitchen godoc
// @Summary Get orders for kitchen
// @Description Get orders for kitchen with priority ordering (Ready > In Progress > Received) and oldest first. Completed orders are excluded.
// @Tags orders
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/kitchen [get]
func (ctrl *OrderController) GetOrdersForKitchen(c *gin.Context) {
	orders, err := ctrl.orderUseCase.GetOrdersForKitchen()
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentOrders(orders)
	c.JSON(http.StatusOK, response)
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an existing order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body dto.UpdateOrderStatusRequest true "status"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/{id}/status [put]
func (ctrl *OrderController) UpdateOrderStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var request dto.UpdateOrderStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	order, err := ctrl.orderUseCase.UpdateOrderStatus(id, &request)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentOrder(order)
	c.JSON(http.StatusOK, response)
}

// DeleteOrder godoc
// @Summary Delete order
// @Description Delete order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /orders/{id} [delete]
func (ctrl *OrderController) DeleteOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ctrl.orderUseCase.DeleteOrder(id)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentSuccess("Order deleted successfully")
	c.JSON(http.StatusOK, response)
}
