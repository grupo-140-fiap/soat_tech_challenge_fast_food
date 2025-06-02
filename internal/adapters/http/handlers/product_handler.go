package handlers

import (
	"net/http"
	"strconv"

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

// GetProductById godoc
// @Summary      Get product by ID
// @Description  Retrieves a product by its unique identifier.
// @Tags         products
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  dto.ProductDTO
// @Router       /products/{id} [get]
func (h *ProductHandler) GetProductById(c *gin.Context) {
	id := c.Param("id")

	product, err := h.productService.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find product",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductByCategory godoc
// @Summary      Get products by category
// @Description  Retrieves a list of products filtered by the specified category.
// @Tags         products
// @Param        category  path      string  true  "Product Category"
// @Success      200       {array}   entities.Product
// @Router       /products/category/{category} [get]
func (h *ProductHandler) GetProductByCategory(c *gin.Context) {
	category := c.Param("category")

	products, err := h.productService.GetProductByCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find product",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, products)
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Creates a new product using the provided JSON payload
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.ProductDTO  true  "Product data"
// @Success      200      {object}  map[string]interface{}  "Product created successfully"
// @Router       /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO

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

// UpdateProduct godoc
// @Summary      Update an existing product
// @Description  Updates the details of an existing product based on the provided JSON payload.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      string  true  "Product ID"
// @Param        product  body      dto.ProductDTO  true  "Product data"
// @Success      200      {object}  map[string]interface{}  "Product updated successfully"
// @Router       /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	productIdStr := c.Param("id")

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Product ID format",
			"error":   "Product ID must be a valid integer",
		})
		return
	}

	var productDTO dto.ProductDTO

	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	err = h.productService.UpdateProduct(productId, &productDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update product",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
	})
}

// DeleteProductById godoc
// @Summary      Delete a product by ID
// @Description  Deletes a product from the database using its unique identifier.
// @Tags         products
// @Param        id   path      string  true  "Product ID"
// @Success      200  {object}  map[string]interface{}  "Product deleted successfully"
// @Router       /products/{id} [delete]
func (h *ProductHandler) DeleteProductById(c *gin.Context) {
	id := c.Param("id")

	err := h.productService.DeleteProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed on find product",
			"error":   err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
