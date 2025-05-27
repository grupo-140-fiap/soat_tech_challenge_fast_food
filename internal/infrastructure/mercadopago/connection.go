package mercadopago

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func NewConnection() (order.Client, error) {
	//err := godotenv.Load()
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Could not load the .env file, using system environment variables")
	}

	// Get the token from the environment variables
	accessToken := os.Getenv("ACCESSTOKEN")
	if accessToken == "" {
		log.Fatal("ACCESSTOKEN not found in .env")
	}

	c, err := config.New(accessToken)
	if err != nil {
		log.Fatalf("Erro ao conectar pagamento: %v", err)
	}

	payClient := order.NewClient(c)

	log.Println("Database connection successfully established")

	return payClient, nil
}
