package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/usecases"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"
)

type CustomerController struct {
	customerUseCase usecases.CustomerUseCase
	presenter       presenters.CustomerPresenter
}

func NewCustomerController(
	customerUseCase usecases.CustomerUseCase,
	presenter presenters.CustomerPresenter,
) *CustomerController {
	return &CustomerController{
		customerUseCase: customerUseCase,
		presenter:       presenter,
	}
}

// CreateCustomer godoc
// @Summary Create new customer
// @Description Create new customer
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body dto.CreateCustomerRequest true "customer"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /customers [post]
func (ctrl *CustomerController) CreateCustomer(c *gin.Context) {
	var request dto.CreateCustomerRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	customer, err := ctrl.customerUseCase.CreateCustomer(&request)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentCustomer(customer)
	c.JSON(http.StatusCreated, response)
}

// GetCustomerByCPF godoc
// @Summary Get customer by CPF
// @Description Get customer by CPF
// @Tags customers
// @Produce json
// @Param cpf path string true "Customer CPF"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /customers/{cpf} [get]
func (ctrl *CustomerController) GetCustomerByCPF(c *gin.Context) {
	cpf := c.Param("cpf")

	customer, err := ctrl.customerUseCase.GetCustomerByCPF(cpf)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := ctrl.presenter.PresentCustomer(customer)
	c.JSON(http.StatusOK, response)
}

// GetCustomerByID godoc
// @Summary Get customer by ID
// @Description Get customer by ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /customers/id/{id} [get]
func (ctrl *CustomerController) GetCustomerByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	customer, err := ctrl.customerUseCase.GetCustomerByID(id)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := ctrl.presenter.PresentCustomer(customer)
	c.JSON(http.StatusOK, response)
}

// UpdateCustomer godoc
// @Summary Update customer
// @Description Update customer information
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body dto.UpdateCustomerRequest true "customer"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /customers/{id} [put]
func (ctrl *CustomerController) UpdateCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var request dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	customer, err := ctrl.customerUseCase.UpdateCustomer(id, &request)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentCustomer(customer)
	c.JSON(http.StatusOK, response)
}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Delete customer by ID
// @Tags customers
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /customers/{id} [delete]
func (ctrl *CustomerController) DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ctrl.customerUseCase.DeleteCustomer(id)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentSuccess("Customer deleted successfully")
	c.JSON(http.StatusOK, response)
}
