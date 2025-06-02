package services

import (
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type AdminService struct {
	orderRepository repositories.OrderRepository
}

func NewAdminService(orderRepository repositories.OrderRepository) *AdminService {
	return &AdminService{
		orderRepository: orderRepository,
	}
}

func (s *AdminService) GetActiveOrders() (*[]entities.Order, error) {
	orders, err := s.orderRepository.GetActiveOrders()
	if err != nil {
		return nil, err
	}

	if orders == nil {
		emptyOrders := make([]entities.Order, 0)
		return &emptyOrders, nil
	}

	return orders, nil
}
