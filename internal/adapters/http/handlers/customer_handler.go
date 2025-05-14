package handlers

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var customerDTO dto.CreateCustomerDTO

	if err := c.ShouldBindJSON(&customerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})

		return
	}

	err := h.customerService.CreateCustomer(&customerDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create customer",
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
	})
}

func (h *CustomerHandler) GetCustomerByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	user, err := h.customerService.GetCustomerByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find customer",
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
