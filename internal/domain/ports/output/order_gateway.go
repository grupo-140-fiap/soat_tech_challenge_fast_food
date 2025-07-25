package output

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"

// OrderGateway defines the contract for order data access operations
type OrderGateway interface {
	Create(order *entities.Order) error
	GetByID(id uint64) (*entities.Order, error)
	GetByCPF(cpf string) ([]*entities.Order, error)
	GetByCustomerID(customerID uint64) ([]*entities.Order, error)
	GetAll() ([]*entities.Order, error)
	GetPendingOrdersForKitchen() ([]*entities.Order, error)
	Update(order *entities.Order) error
	Delete(id uint64) error
}

// OrderItemGateway defines the contract for order item data access operations
type OrderItemGateway interface {
	Create(orderItem *entities.OrderItem) error
	GetByOrderID(orderID uint64) ([]*entities.OrderItem, error)
	Update(orderItem *entities.OrderItem) error
	Delete(id uint64) error
}
