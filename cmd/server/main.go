package main

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/http/router"
)

// @title           SOAT Tech Challenge Fast Food API
// @version         1.0
// @description     API para gerenciamento de pedidos e produtos de lanchonete do SOAT Tech Challenge.

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	router := router.SetupRouter()

  	router.Run(":8080")
}