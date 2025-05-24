package services

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"

type AdminService interface {
	GetActiveOrders() (*[]entities.Order, error)
}
