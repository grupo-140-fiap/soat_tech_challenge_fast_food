package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/usecases"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/persistance/gateways"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/controllers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"
)

type RouterConfig struct {
	Engine *gin.Engine
	DB     *sql.DB
}

func SetupRoutes(config RouterConfig) {
	customerGateway := gateways.NewCustomerGateway(config.DB)
	productGateway := gateways.NewProductGateway(config.DB)
	orderGateway := gateways.NewOrderGateway(config.DB)
	orderItemGateway := gateways.NewOrderItemGateway(config.DB)

	customerUseCase := usecases.NewCustomerUseCase(customerGateway)
	productUseCase := usecases.NewProductUseCase(productGateway)
	orderUseCase := usecases.NewOrderUseCase(orderGateway, orderItemGateway, productGateway)

	customerPresenter := presenters.NewCustomerPresenter()
	productPresenter := presenters.NewProductPresenter()
	orderPresenter := presenters.NewOrderPresenter()

	customerController := controllers.NewCustomerController(customerUseCase, customerPresenter)
	productController := controllers.NewProductController(productUseCase, productPresenter)
	orderController := controllers.NewOrderController(orderUseCase, orderPresenter)

	api := config.Engine.Group("/api/v1")
	{
		customers := api.Group("/customers")
		{
			customers.POST("", customerController.CreateCustomer)
			customers.GET("/:cpf", customerController.GetCustomerByCPF)
			customers.GET("/id/:id", customerController.GetCustomerByID)
			customers.PUT("/:id", customerController.UpdateCustomer)
			customers.DELETE("/:id", customerController.DeleteCustomer)
		}

		products := api.Group("/products")
		{
			products.POST("", productController.CreateProduct)
			products.GET("", productController.GetAllProducts)
			products.GET("/:id", productController.GetProductByID)
			products.GET("/category/:category", productController.GetProductsByCategory)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}

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

	config.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Fast Food API is running",
		})
	})
}
