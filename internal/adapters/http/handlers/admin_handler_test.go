package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAdminService struct {
	mock.Mock
}

func (m *MockAdminService) GetActiveOrders() (*[]entities.Order, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.Order), args.Error(1)
}

func TestNewAdminHandler(t *testing.T) {
	mockService := &MockAdminService{}
	handler := NewAdminHandler(mockService)

	assert.NotNil(t, handler)
	assert.Equal(t, mockService, handler.adminService)
}

func TestAdminHandler_GetActiveOrders_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrders := &[]entities.Order{
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
					ProductID: 200,
					Quantity:  1,
					Price:     25.50,
				},
			},
		},
	}

	mockService := &MockAdminService{}
	mockService.On("GetActiveOrders").Return(mockOrders, nil)

	handler := NewAdminHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin/orders/active", nil)

	handler.GetActiveOrders(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseOrders []entities.Order
	err := json.Unmarshal(w.Body.Bytes(), &responseOrders)
	assert.NoError(t, err)
	assert.Len(t, responseOrders, 2)
	assert.Equal(t, uint64(1), responseOrders[0].ID)
	assert.Equal(t, "preparing", responseOrders[0].Status)
	assert.Equal(t, "123.456.789-00", responseOrders[0].CPF)
	assert.Equal(t, uint64(2), responseOrders[1].ID)
	assert.Equal(t, "ready", responseOrders[1].Status)
	assert.Equal(t, "987.654.321-00", responseOrders[1].CPF)

	mockService.AssertExpectations(t)
}

func TestAdminHandler_GetActiveOrders_EmptyList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrders := &[]entities.Order{}

	mockService := &MockAdminService{}
	mockService.On("GetActiveOrders").Return(mockOrders, nil)

	handler := NewAdminHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin/orders/active", nil)

	handler.GetActiveOrders(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseOrders []entities.Order
	err := json.Unmarshal(w.Body.Bytes(), &responseOrders)
	assert.NoError(t, err)
	assert.Empty(t, responseOrders)

	mockService.AssertExpectations(t)
}

func TestAdminHandler_GetActiveOrders_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedError := errors.New("database connection failed")

	mockService := &MockAdminService{}
	mockService.On("GetActiveOrders").Return(nil, expectedError)

	handler := NewAdminHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin/orders/active", nil)

	handler.GetActiveOrders(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Failed to retrieve active orders", response["message"])
	assert.Equal(t, "database connection failed", response["error"])

	mockService.AssertExpectations(t)
}

func TestAdminHandler_GetActiveOrders_ServiceReturnsNil(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockAdminService{}
	mockService.On("GetActiveOrders").Return(nil, nil)

	handler := NewAdminHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/admin/orders/active", nil)

	handler.GetActiveOrders(c)

	assert.Equal(t, http.StatusOK, w.Code)

	body := w.Body.String()
	assert.Equal(t, "null", body)

	mockService.AssertExpectations(t)
}
