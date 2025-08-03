package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/docs"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/usecases"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/persistance/gateways"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/controllers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interface/presenters"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	paymentGateway := gateways.NewPaymentGateway(config.DB)

	customerUseCase := usecases.NewCustomerUseCase(customerGateway)
	productUseCase := usecases.NewProductUseCase(productGateway)
	orderUseCase := usecases.NewOrderUseCase(orderGateway, orderItemGateway, productGateway, paymentGateway)
	paymentUseCase := usecases.NewPaymentUseCase(paymentGateway, orderGateway)

	customerPresenter := presenters.NewCustomerPresenter()
	productPresenter := presenters.NewProductPresenter()
	orderPresenter := presenters.NewOrderPresenter()
	paymentPresenter := presenters.NewPaymentPresenter()

	customerController := controllers.NewCustomerController(customerUseCase, customerPresenter)
	productController := controllers.NewProductController(productUseCase, productPresenter)
	orderController := controllers.NewOrderController(orderUseCase, orderPresenter)
	paymentController := controllers.NewPaymentController(paymentUseCase, paymentPresenter)

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
			orders.GET("/kitchen", orderController.GetOrdersForKitchen)
			orders.GET("/cpf/:cpf", orderController.GetOrdersByCPF)
			orders.GET("/customer/:customerId", orderController.GetOrdersByCustomerID)
			orders.GET("/:id", orderController.GetOrderByID)
			orders.PUT("/:id/status", orderController.UpdateOrderStatus)
			orders.DELETE("/:id", orderController.DeleteOrder)
		}

		payments := api.Group("/payments")
		{
			payments.POST("", paymentController.CreatePayment)
			payments.GET("/status/:order_id", paymentController.GetPaymentStatus)
			payments.POST("/webhook", paymentController.PaymentWebhook)
		}
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	config.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	config.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Fast Food API is running",
		})
	})
}
