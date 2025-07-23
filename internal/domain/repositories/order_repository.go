package repositories

import "github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"

// OrderRepository defines the contract for order data persistence
type OrderRepository interface {
	Create(order *entities.Order) error
	GetByID(id uint64) (*entities.Order, error)
	GetByCPF(cpf string) ([]*entities.Order, error)
	GetByCustomerID(customerID uint64) ([]*entities.Order, error)
	GetAll() ([]*entities.Order, error)
	Update(order *entities.Order) error
	Delete(id uint64) error
}

// OrderItemRepository defines the contract for order item data persistence
type OrderItemRepository interface {
	Create(orderItem *entities.OrderItem) error
	GetByOrderID(orderID uint64) ([]*entities.OrderItem, error)
	Update(orderItem *entities.OrderItem) error
	Delete(id uint64) error
}
