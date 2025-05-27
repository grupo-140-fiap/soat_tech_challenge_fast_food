package mercadopago

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func NewConnection() (order.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Obtém o token das variáveis de ambiente (pode vir do .env ou do sistema)
	accessToken := os.Getenv("ACCESSTOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("ACCESSTOKEN não encontrado nas variáveis de ambiente")
	}

	c, err := config.New(accessToken)
	if err != nil {
		return nil, fmt.Errorf("erro ao configurar cliente do MercadoPago: %w", err)
	}

	payClient := order.NewClient(c)

	log.Println("Conexão com o MercadoPago estabelecida com sucesso")

	return payClient, nil
}
