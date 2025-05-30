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

func TestNewCustomerRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewCustomerRepository(db)

	assert.NotNil(t, repo)
	assert.IsType(t, &CustomerRepository{}, repo)
}

func TestCustomerRepository_CreateCustomer_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao.silva@email.com",
	}

	expectedQuery := "INSERT INTO customers \\(first_name, last_name, cpf, email\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	mock.ExpectExec(expectedQuery).
		WithArgs(customerDTO.FirstName, customerDTO.LastName, customerDTO.CPF, customerDTO.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewCustomerRepository(db)

	err = repo.CreateCustomer(customerDTO)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_CreateCustomer_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao.silva@email.com",
	}

	expectedQuery := "INSERT INTO customers \\(first_name, last_name, cpf, email\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	expectedError := errors.New("database connection failed")

	mock.ExpectExec(expectedQuery).
		WithArgs(customerDTO.FirstName, customerDTO.LastName, customerDTO.CPF, customerDTO.Email).
		WillReturnError(expectedError)

	repo := NewCustomerRepository(db)

	err = repo.CreateCustomer(customerDTO)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_CreateCustomer_DuplicateCPF(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	customerDTO := &dto.CreateCustomerDTO{
		FirstName: "João",
		LastName:  "Silva",
		CPF:       "123.456.789-00",
		Email:     "joao.silva@email.com",
	}

	expectedQuery := "INSERT INTO customers \\(first_name, last_name, cpf, email\\) VALUES \\(\\?, \\?, \\?, \\?\\)"
	expectedError := errors.New("UNIQUE constraint failed: customers.cpf")

	mock.ExpectExec(expectedQuery).
		WithArgs(customerDTO.FirstName, customerDTO.LastName, customerDTO.CPF, customerDTO.Email).
		WillReturnError(expectedError)

	repo := NewCustomerRepository(db)

	err = repo.CreateCustomer(customerDTO)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE constraint failed")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_GetCustomerByCpf_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cpf := "123.456.789-00"
	expectedCustomer := &entities.Customer{
		ID:        1,
		FirstName: "João",
		LastName:  "Silva",
		CPF:       cpf,
		Email:     "joao.silva@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	expectedQuery := "SELECT id, first_name, last_name, cpf, email FROM customers WHERE cpf = \\?"
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "cpf", "email"}).
		AddRow(expectedCustomer.ID, expectedCustomer.FirstName, expectedCustomer.LastName, expectedCustomer.CPF, expectedCustomer.Email)

	mock.ExpectQuery(expectedQuery).
		WithArgs(cpf).
		WillReturnRows(rows)

	repo := NewCustomerRepository(db)

	customer, err := repo.GetCustomerByCpf(cpf)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, expectedCustomer.ID, customer.ID)
	assert.Equal(t, expectedCustomer.FirstName, customer.FirstName)
	assert.Equal(t, expectedCustomer.LastName, customer.LastName)
	assert.Equal(t, expectedCustomer.CPF, customer.CPF)
	assert.Equal(t, expectedCustomer.Email, customer.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_GetCustomerByCpf_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cpf := "999.999.999-99"
	expectedQuery := "SELECT id, first_name, last_name, cpf, email FROM customers WHERE cpf = \\?"

	mock.ExpectQuery(expectedQuery).
		WithArgs(cpf).
		WillReturnError(sql.ErrNoRows)

	repo := NewCustomerRepository(db)

	customer, err := repo.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Contains(t, err.Error(), "customer with CPF 999.999.999-99 not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_GetCustomerByCpf_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cpf := "123.456.789-00"
	expectedQuery := "SELECT id, first_name, last_name, cpf, email FROM customers WHERE cpf = \\?"
	expectedError := errors.New("database connection failed")

	mock.ExpectQuery(expectedQuery).
		WithArgs(cpf).
		WillReturnError(expectedError)

	repo := NewCustomerRepository(db)

	customer, err := repo.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_GetCustomerByCpf_ScanError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cpf := "123.456.789-00"
	expectedQuery := "SELECT id, first_name, last_name, cpf, email FROM customers WHERE cpf = \\?"

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "cpf", "email"}).
		AddRow("invalid_id", "João", "Silva", cpf, "joao.silva@email.com")

	mock.ExpectQuery(expectedQuery).
		WithArgs(cpf).
		WillReturnRows(rows)

	repo := NewCustomerRepository(db)

	customer, err := repo.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_GetCustomerByCpf_EmptyCPF(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	cpf := ""
	expectedQuery := "SELECT id, first_name, last_name, cpf, email FROM customers WHERE cpf = \\?"

	mock.ExpectQuery(expectedQuery).
		WithArgs(cpf).
		WillReturnError(sql.ErrNoRows)

	repo := NewCustomerRepository(db)

	customer, err := repo.GetCustomerByCpf(cpf)

	assert.Error(t, err)
	assert.Nil(t, customer)
	assert.Contains(t, err.Error(), "customer with CPF  not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCustomerRepository_CreateCustomer_NilCustomer(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewCustomerRepository(db)

	assert.Panics(t, func() {
		repo.CreateCustomer(nil)
	})
}
