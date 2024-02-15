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
	query := `SELECT id, customer_id, product_id, quantity, order_date, order_type, is_seasonal FROM order_items WHERE order_id = $1`
	rows, err := ods.db.QueryContext(ctx, query, orderID)

	if err != nil {
		return orderItems, err
	}

	for rows.Next() {
		var item repository.OrderItem
		err := rows.Scan(&item.ID, &item.Customer_ID, &item.Product_ID, &item.Quantity, &item.Order_Date, &item.Order_type, &item.Is_seasonal)

		if err != nil {
			return []repository.OrderItem{}, err
		}

		orderItems = append(orderItems, item)
	}

	return orderItems, nil
}

// StoreOrderItems stores a batch of order items
func (ods *orderItemStore) StoreOrderItems(ctx context.Context, orderItems []repository.OrderItem) error {
	query := `INSERT INTO order_items (order_id, product_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	for _, orderItem := range orderItems {
		_, err := ods.db.ExecContext(ctx, query, orderItem.ID, orderItem.Product_ID, orderItem.Quantity, orderItem.Is_seasonal)
		if err != nil {
			return err
		}
	}
	return nil
}
