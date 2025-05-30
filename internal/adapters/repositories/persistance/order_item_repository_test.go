package persistance

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderItemRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewOrderItemRepository(db)

	assert.NotNil(t, repo)
	assert.IsType(t, &OrderItemRepository{}, repo)
}

func TestOrderItemRepository_CreateOrderItem_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	expectedError := errors.New("database connection failed")

	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnError(expectedError)

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_ForeignKeyConstraintError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   999,
		ProductId: 200,
		Quantity:  2,
		Price:     19.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	expectedError := errors.New("FOREIGN KEY constraint failed")

	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnError(expectedError)

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FOREIGN KEY constraint failed")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_InvalidProductId(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 999,
		Quantity:  2,
		Price:     19.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	expectedError := errors.New("FOREIGN KEY constraint failed: products.id")

	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnError(expectedError)

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "FOREIGN KEY constraint failed")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_ZeroQuantity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  0,
		Price:     19.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_NegativePrice(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  2,
		Price:     -5.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_LargeQuantity(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	orderItemDTO := &dto.OrderItemDTO{
		ID:        1,
		OrderID:   100,
		ProductId: 200,
		Quantity:  999999,
		Price:     19.99,
	}

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(expectedQuery).
		WithArgs(orderItemDTO.OrderID, orderItemDTO.ProductId, orderItemDTO.Quantity, orderItemDTO.Price).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewOrderItemRepository(db)

	err = repo.CreateOrderItem(orderItemDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_MultipleItems(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

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

	expectedQuery := "INSERT INTO order_items \\(order_id, product_id, quantity, price\\) VALUES \\(\\?, \\?, \\?, \\?\\)"

	for _, item := range orderItems {
		mock.ExpectExec(expectedQuery).
			WithArgs(item.OrderID, item.ProductId, item.Quantity, item.Price).
			WillReturnResult(sqlmock.NewResult(int64(item.ID), 1))
	}

	repo := NewOrderItemRepository(db)

	for _, item := range orderItems {
		err = repo.CreateOrderItem(item)
		assert.NoError(t, err)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestOrderItemRepository_CreateOrderItem_NilOrderItem(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewOrderItemRepository(db)

	assert.Panics(t, func() {
		repo.CreateOrderItem(nil)
	})
}
