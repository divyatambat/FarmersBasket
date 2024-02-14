package dto

import (
	"errors"
	"time"
)

type Order struct {
	ID              int64     `json:"id"`
	CustomerID      int64     `json:"customer_id"`
	ProductID       int64     `json:"product_id"`
	OrderDate       time.Time `json:"order_date"`
	TotalAmount     float64   `json:"total_amount"`
	PaymentOption   string    `json:"payment_option"`
	PaymentStatus   string    `json:"payment_status"`
	OrderStatus     string    `json:"order_status"`
	DeliveryAddress string    `json:"delivery_address"`
	OrderType       string    `json:"order_type"`
	Quantity        int64     `json:"quantity"`
}

type CreateOrderRequest struct {
	CustomerID      int64     `json:"customer_id"`
	ProductID       int64     `json:"product_id"`
	OrderDate       time.Time `json:"order_date"`
	TotalAmount     float64   `json:"total_amount"`
	PaymentOption   string    `json:"payment_option"`
	PaymentStatus   string    `json:"payment_status"`
	OrderStatus     string    `json:"order_status"`
	DeliveryAddress string    `json:"delivery_address"`
	OrderType       string    `json:"order_type"`
	Quantity        int64     `json:"quantity"`
	Product         []Product `json:"products"`
}

type UpdateOrderStatusRequest struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
}

type ProductInfo struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int64  `json:"price"`
	IsSeasonal  bool   `json:"is_seasonal"`
	Quantity    int64  `json:"quantity"`
}

func (req *CreateOrderRequest) Validate() error {
	if req.CustomerID == 0 {
		return errors.New("customer_id cannot be empty")
	}

	if req.ProductID == 0 {
		return errors.New("product_id cannot be empty")
	}

	if req.OrderDate.IsZero() {
		return errors.New("order_date cannot be empty")
	}

	if req.PaymentOption == "" {
		return errors.New("payment_option cannot be empty")
	}

	if req.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	return nil
}

func (req *UpdateOrderStatusRequest) Validate() error {
	if req.OrderID == 0 {
		return errors.New("order_id cannot be empty")
	}

	if req.Status == "" {
		return errors.New("status cannot be empty")
	}

	return nil
}
