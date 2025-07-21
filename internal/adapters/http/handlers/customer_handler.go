package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: customerService,
	}
}

// CreateCustomer godoc
// @Summary Create new customer
// @Description Create new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body dto.CreateCustomerDTO true "customer"
// @Success 200
// @Router /customers [post]
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
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
	})
}

// GetCustomerByCpf godoc
// @Summary      Get customer by CPF
// @Description  Retrieves a customer by their CPF (Cadastro de Pessoas Físicas).
// @Tags         customers
// @Param        cpf   path      string  true  "Customer CPF"
// @Produce      json
// @Success      200  {object}  entities.Customer
// @Router       /customers/{cpf} [get]
func (h *CustomerHandler) GetCustomerByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	user, err := h.customerService.GetCustomerByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find customer",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
