package router

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/http/handlers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/services"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/database/mysql"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/repositories/persistance"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	db, err := mysql.NewConnection()
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

	customerRepository := persistance.NewCustomerRepository(db)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(customerService)

	v1 := router.Group("/api/v1")
	v1.POST("/customers", customerHandler.CreateCustomer)

	return router
}