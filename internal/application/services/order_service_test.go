package services

import (
	"testing"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
)

// Mock implementations for testing
type mockOrderRepository struct{}

func (m *mockOrderRepository) GetOrders() ([]entities.Order, error) {
	return []entities.Order{}, nil
}

func (m *mockOrderRepository) CreateOrder(order *dto.OrderDTO) error {
	return nil
}

func (m *mockOrderRepository) GetOrderById(id string) (*entities.Order, error) {
	return &entities.Order{}, nil
}

func (m *mockOrderRepository) UpdateOrderStatus(id string, status string) error {
	return nil
}

type mockProductService struct{}

func (m *mockProductService) GetProducts() ([]entities.Product, error) {
	return []entities.Product{}, nil
}

func (m *mockProductService) CreateProduct(product *dto.ProductDTO) error {
	return nil
}

func (m *mockProductService) GetProductById(id string) (*entities.Product, error) {
	return &entities.Product{
		ID:    1,
		Name:  "Test Product",
		Price: 19.99,
	}, nil
}

func (m *mockProductService) UpdateProduct(id string, product *dto.ProductDTO) error {
	return nil
}

func (m *mockProductService) DeleteProductById(id string) error {
	return nil
}

func TestOrderService_CreateOrder(t *testing.T) {
	mockOrderRepo := &mockOrderRepository{}
	mockProductSvc := &mockProductService{}

	orderService := NewOrderService(mockOrderRepo, mockProductSvc)

	orderDTO := &dto.OrderDTO{
		CustomerId: 1,
		CPF:        "123.456.789-10",
		Items: []dto.OrderItemDTO{
			{
				ProductId: 1,
				Quantity:  2,
			},
		},
	}

	err := orderService.CreateOrder(orderDTO)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
