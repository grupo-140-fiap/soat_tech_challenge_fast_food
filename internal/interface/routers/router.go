package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/usecases"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/persistance/gateways"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/controllers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"
)

// RouterConfig holds the configuration for setting up routes
type RouterConfig struct {
	Engine *gin.Engine
	DB     *sql.DB
}

// SetupRoutes configures all routes and dependencies following Clean Architecture
// This is the composition root where all dependencies are injected
func SetupRoutes(config RouterConfig) {
	// Initialize Gateways (Infrastructure layer)
	customerGateway := gateways.NewCustomerGateway(config.DB)
	productGateway := gateways.NewProductGateway(config.DB)
	orderGateway := gateways.NewOrderGateway(config.DB)
	orderItemGateway := gateways.NewOrderItemGateway(config.DB)

	// Initialize Use Cases (Application layer)
	customerUseCase := usecases.NewCustomerUseCase(customerGateway)
	productUseCase := usecases.NewProductUseCase(productGateway)
	orderUseCase := usecases.NewOrderUseCase(orderGateway, orderItemGateway, productGateway)

	// Initialize Presenters (Interface layer)
	customerPresenter := presenters.NewCustomerPresenter()
	productPresenter := presenters.NewProductPresenter()
	orderPresenter := presenters.NewOrderPresenter()

	// Initialize Controllers (Interface layer)
	customerController := controllers.NewCustomerController(customerUseCase, customerPresenter)
	productController := controllers.NewProductController(productUseCase, productPresenter)
	orderController := controllers.NewOrderController(orderUseCase, orderPresenter)

	// Setup API routes
	api := config.Engine.Group("/api/v1")
	{
		// Customer routes
		customers := api.Group("/customers")
		{
			customers.POST("", customerController.CreateCustomer)
			customers.GET("/:cpf", customerController.GetCustomerByCPF)
			customers.GET("/id/:id", customerController.GetCustomerByID)
			customers.PUT("/:id", customerController.UpdateCustomer)
			customers.DELETE("/:id", customerController.DeleteCustomer)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.POST("", productController.CreateProduct)
			products.GET("", productController.GetAllProducts)
			products.GET("/:id", productController.GetProductByID)
			products.GET("/category/:category", productController.GetProductsByCategory)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}

		// Order routes
		orders := api.Group("/orders")
		{
			orders.POST("", orderController.CreateOrder)
			orders.GET("", orderController.GetAllOrders)
			orders.GET("/:id", orderController.GetOrderByID)
			orders.GET("/cpf/:cpf", orderController.GetOrdersByCPF)
			orders.GET("/customer/:customerId", orderController.GetOrdersByCustomerID)
			orders.PUT("/:id/status", orderController.UpdateOrderStatus)
			orders.DELETE("/:id", orderController.DeleteOrder)
		}
	}

	// Health check route
	config.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Fast Food API is running",
		})
	})
}
