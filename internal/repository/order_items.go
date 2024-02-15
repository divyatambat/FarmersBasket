package repository

import (
	"context"
	"time"
)

// order item
type OrderItemStorer interface {
	GetOrderItemsByOrderID(ctx context.Context, orderID int64) ([]OrderItem, error)
	StoreOrderItems(ctx context.Context, orderItems []OrderItem) error
}

type OrderItem struct {
	ID          int64 `storm:"id,increment"`
	Customer_ID int64
	Product_ID  int64
	Quantity    int64
	Order_Date  time.Time
	Order_type  string
	Is_seasonal string
}
