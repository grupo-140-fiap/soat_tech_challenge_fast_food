package services

import (
	"errors"
	"testing"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) GetOrders() ([]entities.Order, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Order), args.Error(1)
}

func (m *MockOrderRepository) CreateOrder(order *dto.OrderDTO) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrderById(id string) (*entities.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateOrderStatus(id string, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockOrderRepository) GetActiveOrders() (*[]entities.Order, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.Order), args.Error(1)
}

func TestNewAdminService(t *testing.T) {
	mockRepo := &MockOrderRepository{}
	service := NewAdminService(mockRepo)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.orderRepository)
}

func TestAdminService_GetActiveOrders_Success(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	activeOrders := &[]entities.Order{
		{
			ID:         1,
			CustomerId: 123,
			CPF:        "123.456.789-00",
			Status:     "preparing",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Items: []entities.OrderItem{
				{
					ID:        1,
					OrderID:   1,
					ProductID: 100,
					Quantity:  2,
					Price:     15.99,
				},
			},
		},
		{
			ID:         2,
			CustomerId: 456,
			CPF:        "987.654.321-00",
			Status:     "ready",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Items: []entities.OrderItem{
				{
					ID:        2,
					OrderID:   2,
					ProductID: 101,
					Quantity:  1,
					Price:     25.50,
				},
			},
		},
	}

	mockRepo.On("GetActiveOrders").Return(activeOrders, nil)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 2)
	assert.Equal(t, uint64(1), (*result)[0].ID)
	assert.Equal(t, "preparing", (*result)[0].Status)
	assert.Equal(t, uint64(2), (*result)[1].ID)
	assert.Equal(t, "ready", (*result)[1].Status)

	mockRepo.AssertExpectations(t)
}

func TestAdminService_GetActiveOrders_EmptyResult(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	activeOrders := &[]entities.Order{}

	mockRepo.On("GetActiveOrders").Return(activeOrders, nil)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 0)

	mockRepo.AssertExpectations(t)
}

func TestAdminService_GetActiveOrders_RepositoryError(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	expectedError := errors.New("database connection failed")

	mockRepo.On("GetActiveOrders").Return(nil, expectedError)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}

func TestAdminService_GetActiveOrders_NilResult(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	mockRepo.On("GetActiveOrders").Return(nil, nil)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 0)

	mockRepo.AssertExpectations(t)
}

func TestAdminService_GetActiveOrders_SingleOrder(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	activeOrders := &[]entities.Order{
		{
			ID:         1,
			CustomerId: 123,
			CPF:        "123.456.789-00",
			Status:     "preparing",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Items: []entities.OrderItem{
				{
					ID:        1,
					OrderID:   1,
					ProductID: 100,
					Quantity:  3,
					Price:     12.99,
				},
			},
		},
	}

	mockRepo.On("GetActiveOrders").Return(activeOrders, nil)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 1)
	assert.Equal(t, uint64(1), (*result)[0].ID)
	assert.Equal(t, "preparing", (*result)[0].Status)
	assert.Equal(t, "123.456.789-00", (*result)[0].CPF)
	assert.Len(t, (*result)[0].Items, 1)

	mockRepo.AssertExpectations(t)
}

func TestAdminService_GetActiveOrders_DifferentStatuses(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	activeOrders := &[]entities.Order{
		{
			ID:         1,
			CustomerId: 123,
			CPF:        "123.456.789-00",
			Status:     "received",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         2,
			CustomerId: 456,
			CPF:        "987.654.321-00",
			Status:     "preparing",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         3,
			CustomerId: 789,
			CPF:        "111.222.333-44",
			Status:     "ready",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	mockRepo.On("GetActiveOrders").Return(activeOrders, nil)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 3)
	assert.Equal(t, "received", (*result)[0].Status)
	assert.Equal(t, "preparing", (*result)[1].Status)
	assert.Equal(t, "ready", (*result)[2].Status)

	mockRepo.AssertExpectations(t)
}

func TestAdminService_GetActiveOrders_NetworkError(t *testing.T) {
	mockRepo := &MockOrderRepository{}

	expectedError := errors.New("network timeout")

	mockRepo.On("GetActiveOrders").Return(nil, expectedError)

	service := NewAdminService(mockRepo)

	result, err := service.GetActiveOrders()

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "network timeout")

	mockRepo.AssertExpectations(t)
}
