package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/order"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/docs"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/http/handlers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/repositories/persistance"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/services"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	setHealthRouter(router)
	setCustomerRouter(db, router)
	setProductRouter(db, router)
	setOrdersRouter(db, router)
	setPaymentRouter(payClient, router)
	setAdminRouter(db, router)
	setSwagger(router)

	return router
}

func setHealthRouter(router *gin.Engine) {
	healthHandler := handlers.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)
}

func setCustomerRouter(db *sql.DB, router *gin.Engine) {
	//mover pro cmd e receber ao inves de db *sql.DB, a interface do New<metodo>
	customerRepository := persistance.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)

	v1 := router.Group("/api/v1")
	v1.POST("/customers", customerHandler.CreateCustomer)
	v1.GET("/customers/:cpf", customerHandler.GetCustomerByCpf)
}

func setProductRouter(db *sql.DB, router *gin.Engine) {
	//mover pro cmd e receber ao inves de db *sql.DB, a interface do New<metodo>
	productRepository := persistance.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	v1 := router.Group("/api/v1")
	v1.GET("/products/:id", productHandler.GetProductById)
	v1.GET("/products/category/:category", productHandler.GetProductByCategory)

	v1.POST("/products", productHandler.CreateProduct)
	v1.PUT("/products/:id", productHandler.UpdateProduct)
	v1.DELETE("/products/:id", productHandler.DeleteProductById)

}

func setOrdersRouter(db *sql.DB, router *gin.Engine) {
	//mover pro cmd e receber ao inves de db *sql.DB, a interface do New<metodo>
	orderRepository := persistance.NewOrderRepository(db)
	productRepository := persistance.NewProductRepository(db)

	productService := services.NewProductService(productRepository)
	orderService := services.NewOrderService(orderRepository, productService)
	orderHandler := handlers.NewOrderHandler(orderService)

	v1 := router.Group("/api/v1")
	v1.GET("/orders", orderHandler.GetOrders)
	v1.POST("/orders", orderHandler.CreateOrder)
	v1.GET("/orders/:id", orderHandler.GetOrderById)
	v1.PATCH("/orders/:id/status", orderHandler.UpdateOrderStatus)
}

func setPaymentRouter(payClient order.Client, router *gin.Engine) {
	//mover pro cmd e receber ao inves de db *sql.DB, a interface do New<metodo>
	paymentRepository := persistance.NewPaymentRepository(payClient)
	paymentService := services.NewPaymentService(paymentRepository)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	v1 := router.Group("/api/v1")
	v1.POST("/checkout", paymentHandler.CreatePayment)
}

func setAdminRouter(db *sql.DB, router *gin.Engine) {
	//mover pro cmd e receber ao inves de db *sql.DB, a interface do New<metodo>
	orderRepository := persistance.NewOrderRepository(db)
	adminService := services.NewAdminService(orderRepository)
	adminHandler := handlers.NewAdminHandler(adminService)

	v1 := router.Group("/api/v1")
	v1.GET("/admin/orders/active", adminHandler.GetActiveOrders)
}

func setSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
