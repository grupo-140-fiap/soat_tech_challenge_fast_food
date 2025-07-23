package gateways

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/entities"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/repositories"
)

// productGateway implements the ProductRepository interface
// This is the gateway that handles data persistence for products
type productGateway struct {
	db *sql.DB
}

// NewProductGateway creates a new product gateway
func NewProductGateway(db *sql.DB) repositories.ProductRepository {
	return &productGateway{
		db: db,
	}
}

// Create persists a new product
func (g *productGateway) Create(product *entities.Product) error {
	query := `
		INSERT INTO products (name, description, price, category, image_url, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := g.db.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Category,
		product.ImageUrl,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = uint64(id)
	return nil
}

// GetByID retrieves a product by ID
func (g *productGateway) GetByID(id uint64) (*entities.Product, error) {
	query := `
		SELECT id, name, description, price, category, image_url, created_at, updated_at
		FROM products
		WHERE id = ?
	`

	row := g.db.QueryRow(query, id)

	var product entities.Product
	var price string
	var createdAt, updatedAt string

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&price,
		&product.Category,
		&product.ImageUrl,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	// Parse price
	if priceFloat, err := strconv.ParseFloat(price, 32); err == nil {
		product.Price = float32(priceFloat)
	}

	// Parse timestamps
	product.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	product.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return &product, nil
}

// GetByCategory retrieves products by category
func (g *productGateway) GetByCategory(category string) ([]*entities.Product, error) {
	query := `
		SELECT id, name, description, price, category, image_url, created_at, updated_at
		FROM products
		WHERE category = ?
		ORDER BY name
	`

	rows, err := g.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product

	for rows.Next() {
		var product entities.Product
		var price string
		var createdAt, updatedAt string

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&price,
			&product.Category,
			&product.ImageUrl,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			continue
		}

		// Parse price
		if priceFloat, err := strconv.ParseFloat(price, 32); err == nil {
			product.Price = float32(priceFloat)
		}

		// Parse timestamps
		product.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		product.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		products = append(products, &product)
	}

	return products, nil
}

// GetAll retrieves all products
func (g *productGateway) GetAll() ([]*entities.Product, error) {
	query := `
		SELECT id, name, description, price, category, image_url, created_at, updated_at
		FROM products
		ORDER BY category, name
	`

	rows, err := g.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product

	for rows.Next() {
		var product entities.Product
		var price string
		var createdAt, updatedAt string

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&price,
			&product.Category,
			&product.ImageUrl,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			continue
		}

		// Parse price
		if priceFloat, err := strconv.ParseFloat(price, 32); err == nil {
			product.Price = float32(priceFloat)
		}

		// Parse timestamps
		product.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		product.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

		products = append(products, &product)
	}

	return products, nil
}

// Update updates an existing product
func (g *productGateway) Update(product *entities.Product) error {
	query := `
		UPDATE products
		SET name = ?, description = ?, price = ?, category = ?, image_url = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := g.db.Exec(query,
		product.Name,
		product.Description,
		product.Price,
		product.Category,
		product.ImageUrl,
		product.UpdatedAt,
		product.ID,
	)

	return err
}

// Delete removes a product
func (g *productGateway) Delete(id uint64) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := g.db.Exec(query, id)
	return err
}
