package repository

import (
	"context"
	"database/sql"

	"github.com/divyatambat/FarmersBasket/internal/repository"
)

// orderItemStore implementation
type orderItemStore struct {
	db *sql.DB
}

// NewOrderItemRepo creates a new OrderItemStorer instance
func NewOrderItemRepo(db *sql.DB) repository.OrderItemStorer {
	return &orderItemStore{db: db}
}

// GetOrderItemsByOrderID retrieves order items by order ID
func (ods *orderItemStore) GetOrderItemsByOrderID(ctx context.Context, orderID int64) ([]repository.OrderItem, error) {
	var orderItems []repository.OrderItem
	// query := `SELECT * FROM order_items WHERE order_id = $1`
	// err := tx.SelectContext(ctx, &orderItems, query, orderID)
	return orderItems, nil
}

// StoreOrderItems stores a batch of order items
func (ods *orderItemStore) StoreOrderItems(ctx context.Context, orderItems []repository.OrderItem) error {
	// query := `INSERT INTO order_items (order_id, product_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	// for _, orderItem := range orderItems {
	// 	_, err := tx.ExecContext(ctx, query, orderItem.Order_ID, orderItem.Product_ID, orderItem.Quantity, orderItem.Is_seasonal)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
