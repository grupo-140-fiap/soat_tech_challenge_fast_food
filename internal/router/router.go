package router

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interfaces/http/handlers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/services"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/persistance/mysql"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	db, err := mysql.NewConnection()
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

	userRepository := mysql.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	v1 := router.Group("/v1")
	v1.POST("/users", userHandler.CreateUser)

	return router
}