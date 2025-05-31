package services

import (
	"errors"
	"testing"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderItemRepository struct {
	mock.Mock
}

func (m *MockOrderItemRepository) CreateOrderItem(orderItem *dto.OrderItemDTO) error {
	args := m.Called(orderItem)
	return args.Error(0)
}

func TestNewOrderItemService(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}
	service := NewOrderItemService(mockRepo)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.orderItemRepository)
}

func TestOrderItemService_CreateOrderItem_Success(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	mockRepo.On("CreateOrderItem", orderItemDTO).Return(nil)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_RepositoryError(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("CreateOrderItem", orderItemDTO).Return(expectedError)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_ForeignKeyError(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   999,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	expectedError := errors.New("FOREIGN KEY constraint failed")
	mockRepo.On("CreateOrderItem", orderItemDTO).Return(expectedError)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FOREIGN KEY constraint failed")
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_ZeroQuantity(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  0,
		Price:     19.99,
	}

	mockRepo.On("CreateOrderItem", orderItemDTO).Return(nil)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_NegativePrice(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     -5.99,
	}

	mockRepo.On("CreateOrderItem", orderItemDTO).Return(nil)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_LargeQuantity(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  999999,
		Price:     19.99,
	}

	mockRepo.On("CreateOrderItem", orderItemDTO).Return(nil)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_InvalidProductId(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 999,
		Quantity:  2,
		Price:     19.99,
	}

	expectedError := errors.New("product not found")
	mockRepo.On("CreateOrderItem", orderItemDTO).Return(expectedError)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "product not found")
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_NetworkTimeout(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	expectedError := errors.New("network timeout")
	mockRepo.On("CreateOrderItem", orderItemDTO).Return(expectedError)

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "network timeout")
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_NilOrderItem(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	mockRepo.On("CreateOrderItem", (*dto.OrderItemDTO)(nil)).Return(errors.New("invalid order item"))

	service := NewOrderItemService(mockRepo)

	err := service.CreateOrderItem(nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid order item")
	mockRepo.AssertExpectations(t)
}

func TestOrderItemService_CreateOrderItem_MultipleItemsSequence(t *testing.T) {
	mockRepo := &MockOrderItemRepository{}

	orderItems := []*dto.OrderItemDTO{
		{
			ID:        1,
			OrderID:   100,
			ProductId: 200,
			Quantity:  2,
			Price:     19.99,
		},
		{
			ID:        2,
			OrderID:   100,
			ProductId: 201,
			Quantity:  1,
			Price:     15.50,
		},
		{
			ID:        3,
			OrderID:   100,
			ProductId: 202,
			Quantity:  3,
			Price:     8.99,
		},
	}

	for _, item := range orderItems {
		mockRepo.On("CreateOrderItem", item).Return(nil)
	}

	service := NewOrderItemService(mockRepo)

	for _, item := range orderItems {
		err := service.CreateOrderItem(item)
		assert.NoError(t, err)
	}

	mockRepo.AssertExpectations(t)
}
