package services

import (
	"testing"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities/dto"
)

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

func (m *mockOrderRepository) GetActiveOrders() (*[]entities.Order, error) {
	orders := []entities.Order{}
	return &orders, nil
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

func (m *mockProductService) UpdateProduct(id int, product *dto.ProductDTO) error {
	return nil
}

func (m *mockProductService) DeleteProductById(id string) error {
	return nil
}

func (m *mockProductService) GetProductByCategory(category string) ([]entities.Product, error) {
	return []entities.Product{}, nil
}

func TestOrderService_UpdateOrderStatus_ValidStatus(t *testing.T) {
	mockOrderRepo := &mockOrderRepository{}
	mockProductSvc := &mockProductService{}

	orderService := NewOrderService(mockOrderRepo, mockProductSvc)

	validStatuses := []string{"received", "preparation", "ready", "completed"}

	for _, status := range validStatuses {
		err := orderService.UpdateOrderStatus("1", status)
		if err != nil {
			t.Errorf("Expected no error for valid status '%s', got %v", status, err)
		}
	}
}

func TestOrderService_UpdateOrderStatus_InvalidStatus(t *testing.T) {
	mockOrderRepo := &mockOrderRepository{}
	mockProductSvc := &mockProductService{}

	orderService := NewOrderService(mockOrderRepo, mockProductSvc)

	invalidStatuses := []string{"invalid", "pending", "cancelled", ""}

	for _, status := range invalidStatuses {
		err := orderService.UpdateOrderStatus("1", status)
		if err == nil {
			t.Errorf("Expected error for invalid status '%s', got nil", status)
		}
	}
}

func TestIsValidOrderStatus(t *testing.T) {
	tests := []struct {
		status   string
		expected bool
	}{
		{"received", true},
		{"preparation", true},
		{"ready", true},
		{"completed", true},
		{"invalid", false},
		{"pending", false},
		{"cancelled", false},
		{"", false},
	}

	for _, test := range tests {
		result := isValidOrderStatus(test.status)
		if result != test.expected {
			t.Errorf("isValidOrderStatus(%s) = %v; expected %v", test.status, result, test.expected)
		}
	}
}
