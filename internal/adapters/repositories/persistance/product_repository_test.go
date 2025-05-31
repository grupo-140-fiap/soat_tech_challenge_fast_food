package persistance

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewProductRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepository(db)

	assert.NotNil(t, repo)
	assert.IsType(t, &ProductRepository{}, repo)
}

func TestProductRepository_CreateProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productDTO := &dto.ProductDTO{
		Name:        "Burger",
		Description: "Delicious burger",
		Price:       "15.99",
		Category:    "main",
	}

	expectedQuery := "INSERT INTO products \\(name, description, price, category\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(expectedQuery).
		WithArgs(productDTO.Name, productDTO.Description, productDTO.Price, productDTO.Category).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepository(db)

	err = repo.CreateProduct(productDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_CreateProduct_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productDTO := &dto.ProductDTO{
		Name:        "Burger",
		Description: "Delicious burger",
		Price:       "15.99",
		Category:    "main",
	}

	expectedQuery := "INSERT INTO products \\(name, description, price, category\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	expectedError := errors.New("database connection failed")

	mock.ExpectExec(expectedQuery).
		WithArgs(productDTO.Name, productDTO.Description, productDTO.Price, productDTO.Category).
		WillReturnError(expectedError)

	repo := NewProductRepository(db)

	err = repo.CreateProduct(productDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productID := "1"
	expectedProduct := &entities.Product{
		ID:          1,
		Name:        "Burger",
		Description: "Delicious burger",
		Price:       15.99,
		Category:    "main",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category"}).
		AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Description, expectedProduct.Price, expectedProduct.Category)

	mock.ExpectQuery(expectedQuery).
		WithArgs(productID).
		WillReturnRows(rows)

	repo := NewProductRepository(db)

	product, err := repo.GetProductById(productID)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, expectedProduct.ID, product.ID)
	assert.Equal(t, expectedProduct.Name, product.Name)
	assert.Equal(t, expectedProduct.Description, product.Description)
	assert.Equal(t, expectedProduct.Price, product.Price)
	assert.Equal(t, expectedProduct.Category, product.Category)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductById_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productID := "999"
	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE id = \\?"

	mock.ExpectQuery(expectedQuery).
		WithArgs(productID).
		WillReturnError(sql.ErrNoRows)

	repo := NewProductRepository(db)

	product, err := repo.GetProductById(productID)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Contains(t, err.Error(), "product with ID 999 not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductById_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productID := "1"
	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE id = \\?"
	expectedError := errors.New("database connection failed")

	mock.ExpectQuery(expectedQuery).
		WithArgs(productID).
		WillReturnError(expectedError)

	repo := NewProductRepository(db)

	product, err := repo.GetProductById(productID)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductByCategory_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	category := "main"
	expectedProducts := []entities.Product{
		{
			ID:          1,
			Name:        "Burger",
			Description: "Delicious burger",
			Price:       15.99,
			Category:    "main",
		},
		{
			ID:          2,
			Name:        "Pizza",
			Description: "Tasty pizza",
			Price:       25.50,
			Category:    "main",
		},
	}

	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE category = \\?"
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category"}).
		AddRow(expectedProducts[0].ID, expectedProducts[0].Name, expectedProducts[0].Description, expectedProducts[0].Price, expectedProducts[0].Category).
		AddRow(expectedProducts[1].ID, expectedProducts[1].Name, expectedProducts[1].Description, expectedProducts[1].Price, expectedProducts[1].Category)

	mock.ExpectQuery(expectedQuery).
		WithArgs(category).
		WillReturnRows(rows)

	repo := NewProductRepository(db)

	products, err := repo.GetProductByCategory(category)

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 2)
	assert.Equal(t, expectedProducts[0].Name, products[0].Name)
	assert.Equal(t, expectedProducts[1].Name, products[1].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductByCategory_EmptyResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	category := "nonexistent"
	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE category = \\?"
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category"})

	mock.ExpectQuery(expectedQuery).
		WithArgs(category).
		WillReturnRows(rows)

	repo := NewProductRepository(db)

	products, err := repo.GetProductByCategory(category)

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductByCategory_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	category := "main"
	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE category = \\?"
	expectedError := errors.New("products with CATEGORY main not found")

	mock.ExpectQuery(expectedQuery).
		WithArgs(category).
		WillReturnError(expectedError)

	repo := NewProductRepository(db)

	products, err := repo.GetProductByCategory(category)

	assert.Error(t, err)
	assert.Nil(t, products)
	assert.Contains(t, err.Error(), "products with CATEGORY main not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_GetProductByCategory_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	category := "main"
	expectedQuery := "SELECT id, name, description, price, category FROM products WHERE category = \\?"
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category"}).
		AddRow("invalid_id", "Burger", "Delicious burger", 15.99, "main")

	mock.ExpectQuery(expectedQuery).
		WithArgs(category).
		WillReturnRows(rows)

	repo := NewProductRepository(db)

	products, err := repo.GetProductByCategory(category)

	assert.Error(t, err)
	assert.Nil(t, products)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_UpdateProduct_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productDTO := &dto.ProductDTO{
		ID:          1,
		Name:        "Updated Burger",
		Description: "Updated delicious burger",
		Price:       "18.99",
		Category:    "main",
	}

	expectedQuery := "UPDATE products SET name = \\?, description = \\?, price = \\?, category = \\? WHERE id = \\?"
	mock.ExpectExec(expectedQuery).
		WithArgs(productDTO.Name, productDTO.Description, productDTO.Price, productDTO.Category, productDTO.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepository(db)

	err = repo.UpdateProduct(productDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_UpdateProduct_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productDTO := &dto.ProductDTO{
		ID:          1,
		Name:        "Updated Burger",
		Description: "Updated delicious burger",
		Price:       "18.99",
		Category:    "main",
	}

	expectedQuery := "UPDATE products SET name = \\?, description = \\?, price = \\?, category = \\? WHERE id = \\?"
	expectedError := errors.New("database connection failed")

	mock.ExpectExec(expectedQuery).
		WithArgs(productDTO.Name, productDTO.Description, productDTO.Price, productDTO.Category, productDTO.ID).
		WillReturnError(expectedError)

	repo := NewProductRepository(db)

	err = repo.UpdateProduct(productDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_UpdateProduct_NoRowsAffected(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productDTO := &dto.ProductDTO{
		ID:          999,
		Name:        "Updated Burger",
		Description: "Updated delicious burger",
		Price:       "18.99",
		Category:    "main",
	}

	expectedQuery := "UPDATE products SET name = \\?, description = \\?, price = \\?, category = \\? WHERE id = \\?"
	mock.ExpectExec(expectedQuery).
		WithArgs(productDTO.Name, productDTO.Description, productDTO.Price, productDTO.Category, productDTO.ID).
		WillReturnResult(sqlmock.NewResult(1, 0))

	repo := NewProductRepository(db)

	err = repo.UpdateProduct(productDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_DeleteProductById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productID := "1"
	expectedQuery := "DELETE FROM products WHERE id = \\?"
	mock.ExpectExec(expectedQuery).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewProductRepository(db)

	err = repo.DeleteProductById(productID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_DeleteProductById_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productID := "1"
	expectedQuery := "DELETE FROM products WHERE id = \\?"
	expectedError := errors.New("database connection failed")

	mock.ExpectExec(expectedQuery).
		WithArgs(productID).
		WillReturnError(expectedError)

	repo := NewProductRepository(db)

	err = repo.DeleteProductById(productID)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_DeleteProductById_NoRowsAffected(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	productID := "999"
	expectedQuery := "DELETE FROM products WHERE id = \\?"
	mock.ExpectExec(expectedQuery).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 0))

	repo := NewProductRepository(db)

	err = repo.DeleteProductById(productID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProductRepository_CreateProduct_NilProduct(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepository(db)

	assert.Panics(t, func() {
		repo.CreateProduct(nil)
	})
}

func TestProductRepository_UpdateProduct_NilProduct(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewProductRepository(db)

	assert.Panics(t, func() {
		repo.UpdateProduct(nil)
	})
}
