package repository

import (
	"context"
	"time"
)

// order
type OrderStorer interface {
	GetOrderByID(ctx context.Context, orderID int64) (Order, error)
	CreateOrder(ctx context.Context, order Order) (Order, error)
	UpdateOrderStatus(ctx context.Context, orderID int64, status string) error
	// UpdateOrderDispatchDate(ctx context.Context, orderID int64, dispatchedAt time.Time) error
	ListOrders(ctx context.Context) ([]Order, error)
}

type Order struct {
	ID               uint
	Customer_ID      int64
	Product_ID       int64
	Order_Date       time.Time
	Total_Amount     float64
	Payment_option   string
	Payment_status   string
	Order_Status     string
	Order_type       string
	Delivery_Address string
	Quantity         float64
}
