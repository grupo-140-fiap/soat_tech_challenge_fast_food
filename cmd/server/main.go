package main

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/adapters/http/router"
)

func main() {
	router := router.SetupRouter()

  	router.Run()
}