package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/usecases"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"
)

// ProductController handles HTTP requests for product operations
type ProductController struct {
	productUseCase usecases.ProductUseCase
	presenter      presenters.ProductPresenter
}

// NewProductController creates a new product controller
func NewProductController(
	productUseCase usecases.ProductUseCase,
	presenter presenters.ProductPresenter,
) *ProductController {
	return &ProductController{
		productUseCase: productUseCase,
		presenter:      presenter,
	}
}

// CreateProduct godoc
// @Summary Create new product
// @Description Create new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "product"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [post]
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var request dto.CreateProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := ctrl.productUseCase.CreateProduct(&request)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentProduct(product)
	c.JSON(http.StatusCreated, response)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [get]
func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := ctrl.productUseCase.GetProductByID(id)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := ctrl.presenter.PresentProduct(product)
	c.JSON(http.StatusOK, response)
}

// GetProductsByCategory godoc
// @Summary Get products by category
// @Description Get products by category (snack, drink, dessert, side)
// @Tags products
// @Produce json
// @Param category path string true "Product Category"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/category/{category} [get]
func (ctrl *ProductController) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")

	products, err := ctrl.productUseCase.GetProductsByCategory(category)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := ctrl.presenter.PresentProducts(products)
	c.JSON(http.StatusOK, response)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products [get]
func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
	products, err := ctrl.productUseCase.GetAllProducts()
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentProducts(products)
	c.JSON(http.StatusOK, response)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product information
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.UpdateProductRequest true "product"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [put]
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var request dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := ctrl.productUseCase.UpdateProduct(id, &request)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentProduct(product)
	c.JSON(http.StatusOK, response)
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /products/{id} [delete]
func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = ctrl.productUseCase.DeleteProduct(id)
	if err != nil {
		response := ctrl.presenter.PresentError(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := ctrl.presenter.PresentSuccess("Product deleted successfully")
	c.JSON(http.StatusOK, response)
}
