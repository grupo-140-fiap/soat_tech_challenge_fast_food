package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/input/services"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productDTO dto.CreateProductDTO

	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})

		return
	}

	err := h.productService.CreateProduct(&productDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create product",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
	})
}

func (h *ProductHandler) GetProductById(c *gin.Context) {
	id := c.Param("id")

	user, err := h.productService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find product",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, user)
}
