package main

import (
	"log"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/http/router"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/mercadopago"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/repositories/mysql"
)

// @title           SOAT Tech Challenge Fast Food API
// @version         1.0
// @description     API para gerenciamento de pedidos e produtos de lanchonete do SOAT Tech Challenge.

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	//Mover isso para o cmd/main.go
	db, err := mysql.NewConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	payClient, err := mercadopago.NewConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao mercadopago: %v", err)
	}

	router := router.SetupRouter()

	router.Run(":8080")
}
