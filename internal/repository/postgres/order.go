package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/divyatambat/FarmersBasket/internal/repository"
	_ "github.com/lib/pq"
)

type orderStore struct {
	db *sql.DB
}

func NewOrderRepo(db *sql.DB) repository.OrderStorer {
	return &orderStore{db: db}
}

func (os *orderStore) GetOrderByID(ctx context.Context, orderID int64) (repository.Order, error) {
	var order repository.Order

	row := os.db.QueryRowContext(ctx, `
		SELECT o.*, u.name AS customer_name, p.name AS product_name
		FROM orders o
		JOIN users u ON o.customer_id = u.id
		JOIN products p ON o.product_id = p.id
		WHERE o.id = $1
	`, orderID)
	err := row.Scan(&order.ID, &order.Customer_ID, &order.Product_ID, &order.Order_Date, &order.Total_Amount,
		&order.Payment_option, &order.Payment_status, &order.Order_Status, &order.Delivery_Address,
		&order.Order_type, &order.Quantity)

	if err == sql.ErrNoRows {
		return repository.Order{}, errors.New("order not found")
	} else if err != nil {
		return repository.Order{}, fmt.Errorf("error getting order: %w", err)
	}

	return order, nil
}

func (os *orderStore) CreateOrder(ctx context.Context, order repository.Order) (repository.Order, error) {

	_, err := os.db.ExecContext(ctx, `
		INSERT INTO orders (customer_id, product_id, order_date, total_amount, payment_option, payment_status, order_status, delivery_address, order_type, quantity)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id`, order.Customer_ID, order.Product_ID, order.Order_Date, order.Total_Amount, order.Payment_option, order.Payment_status, order.Order_Status, order.Delivery_Address, order.Order_type, order.Quantity)

	if err != nil {
		return repository.Order{}, fmt.Errorf("error creating order: %w", err)
	}

	var insertedOrderID uint
	if err != nil {
		return repository.Order{}, fmt.Errorf("error retrieving inserted order ID: %w", err)
	}

	order.ID = insertedOrderID
	return order, nil
}

func (os *orderStore) UpdateOrderStatus(ctx context.Context, orderID int64, status string) error { // Check if transaction is available for prepared statements

	_, err := os.db.ExecContext(ctx, `
		UPDATE orders SET order_status = 1, updated_at = 2
		WHERE id = 3
	`, status, time.Now(), orderID)
	if err != nil {
		return err
	}
	return nil
}

func (os *orderStore) ListOrders(ctx context.Context) ([]repository.Order, error) {
	orderList := make([]repository.Order, 0)

	rows, err := os.db.QueryContext(ctx, `
        SELECT id, order_date, total_amount, payment_option, payment_status, order_status, order_type, delivery_address, quantity
        FROM orders
    `)
	if err != nil {
		return nil, fmt.Errorf("error listing orders: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order repository.Order
		err := rows.Scan(&order.ID, &order.Order_Date, &order.Total_Amount, &order.Payment_option, &order.Payment_status, &order.Order_Status, &order.Order_type, &order.Delivery_Address, &order.Quantity)
		if err != nil {
			return nil, fmt.Errorf("error scanning order row: %w", err)
		}
		orderList = append(orderList, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating order rows: %w", err)
	}

	return orderList, nil
}
