package pkg

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID       *int64 `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Phone    *int64 `db:"phone_number"`
	UserType string `db:"user_type"`
}

type Product struct {
	ID          int64           `db:"id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Category    string          `db:"category"`
	Price       decimal.Decimal `db:"price"`
	Seasonal    bool            `db:"is_seasonal"`
	Quantity    int             `db:"quantity"`
}

type Order struct {
	ID            int64     `db:"id,pk"`
	CustomerID    int64     `db:"customer_id,ref=user.id"`
	ProductID     int64     `db:"product_id,ref=product.id"`
	OrderDate     time.Time `db:"order_date"`
	TotalAmount   float64   `db:"total_amount"`
	PaymentOption string    `db:"payment_option"`
	PaymentStatus string    `db:"payment_status"`
	OrderStatus   string    `db:"order_status"`
	DeliveryAddr  string    `db:"delivery_address"`
	OrderType     string    `db:"order_type"`
	Quantity      int       `db:"quantity"`
}
