package router

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/http/handlers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/repositories/persistance"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/services"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/database/mysql"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	db, err := mysql.NewConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	setCustomerRouter(db, router)
	setProductRouter(db, router)

	return router
}

func setCustomerRouter(db *sql.DB, router *gin.Engine) {
	customerRepository := persistance.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)

	v1 := router.Group("/api/v1")
	v1.POST("/customers", customerHandler.CreateCustomer)
	v1.GET("/customers/:cpf", customerHandler.GetCustomerByCpf)
}

func setProductRouter(db *sql.DB, router *gin.Engine) {
	productRepository := persistance.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	v1 := router.Group("/api/v1")
	v1.GET("/products/:id", productHandler.GetProductById)
	v1.GET("/products/category/:category", productHandler.GetProductByCategory)
	v1.POST("/products", productHandler.CreateProduct)
	v1.PUT("/products", productHandler.UpdateProduct)
	v1.DELETE("/products/:id", productHandler.DeleteProductById)

}
