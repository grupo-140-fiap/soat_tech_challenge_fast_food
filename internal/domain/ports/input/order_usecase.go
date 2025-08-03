package input

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"

// OrderUseCase defines the contract for order business operations
type OrderUseCase interface {
	CreateOrder(request *dto.CreateOrderRequest) (*dto.OrderResponse, error)
	GetOrderByID(id uint64) (*dto.OrderResponse, error)
	GetOrdersByCPF(cpf string) ([]*dto.OrderResponse, error)
	GetOrdersByCustomerID(customerID uint64) ([]*dto.OrderResponse, error)
	GetAllOrders() ([]*dto.OrderResponse, error)
	GetOrdersForKitchen() ([]*dto.OrderResponse, error)
	UpdateOrderStatus(id uint64, request *dto.UpdateOrderStatusRequest) (*dto.OrderResponse, error)
	DeleteOrder(id uint64) error
}
