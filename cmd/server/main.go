package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/interfaces/http/handlers"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/services"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/infrastructure/persistance/mysql"
)

func main() {
    router := gin.Default()

	userRepository := mysql.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	v1 := router.Group("/v1")
	v1.POST("/users", userHandler.CreateUser)

  	router.Run()
}