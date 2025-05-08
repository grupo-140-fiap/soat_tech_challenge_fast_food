package main

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/router"
)

func main() {
	router := router.SetupRouter()

  	router.Run()
}