package persistance

import (
	"database/sql"
	"fmt"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repositories.ProductRepository {
	return &ProductRepository{db: db}
}

func (u *ProductRepository) CreateProduct(product *dto.ProductDTO) error {
	query := "INSERT INTO products (name, description, price, category) VALUES (?, ?, ?, ?)"

	_, err := u.db.Exec(query, product.Name, product.Description, product.Price, product.Category)

	if err != nil {
		return err
	}

	return nil
}

func (u *ProductRepository) GetProductById(id string) (*entities.Product, error) {
	query := "SELECT id, name, description, price, category FROM products WHERE id = ?"
	row := u.db.QueryRow(query, id)

	var product entities.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %s not found", id)
		}

		return nil, err
	}

	return &product, nil
}

func (u *ProductRepository) UpdateProduct(product *dto.ProductDTO) error {
	query := "UPDATE products SET name = ?, description = ?, price = ?, category = ? WHERE id = ?"

	_, err := u.db.Exec(query, product.Name, product.Description, product.Price, product.Category, product.ID)

	if err != nil {
		return err
	}

	return nil
}
