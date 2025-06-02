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
	query := "INSERT INTO products (name, description, price, category, image_url) VALUES (?, ?, ?, ?, ?)"

	_, err := u.db.Exec(query, product.Name, product.Description, product.Price, product.Category, product.ImageUrl)

	if err != nil {
		return err
	}

	return nil
}

func (u *ProductRepository) GetProductById(id string) (*entities.Product, error) {
	query := "SELECT id, name, description, price, category, image_url FROM products WHERE id = ?"
	row := u.db.QueryRow(query, id)

	var product entities.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.ImageUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %s not found", id)
		}

		return nil, err
	}

	return &product, nil
}

func (u *ProductRepository) GetProductByCategory(category string) ([]entities.Product, error) {
	rows, err := u.db.Query("SELECT id, name, description, price, category, image_url FROM products WHERE category = ?", category)
	if err != nil {
		return nil, fmt.Errorf("products with CATEGORY %s not found", category)
	}
	defer rows.Close()

	var products []entities.Product = make([]entities.Product, 0)
	for rows.Next() {
		var p entities.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (u *ProductRepository) UpdateProduct(producId int, product *dto.ProductDTO) error {
	query := "UPDATE products SET name = ?, description = ?, price = ?, category = ?, image_url = ? WHERE id = ?"

	_, err := u.db.Exec(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Category,
		product.ImageUrl,
		producId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *ProductRepository) DeleteProductById(id string) error {
	query := "DELETE FROM products WHERE id = ?"

	_, err := u.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
