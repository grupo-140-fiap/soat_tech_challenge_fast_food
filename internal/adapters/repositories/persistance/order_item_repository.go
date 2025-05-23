package persistance

import (
	"database/sql"

	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/application/dto"
	"github.com/samuellalvs/soat_tech_challenge_fast_food/internal/domain/ports/output/repositories"
)

type OrderItemRepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) repositories.OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (u *OrderItemRepository) CreateOrderItem(item *dto.OrderItemDTO) error {
	query := "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)"

	_, err := u.db.Exec(query, item.OrderId, item.ProductId, item.Quantity, item.Price)

	if err != nil {
		return err
	}

	return nil
}
