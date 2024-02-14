package repository

import "context"

// product
type ProductStorer interface {
	GetProductByID(ctx context.Context, tx Transaction, productID int64) (Product, error)
	ListProducts(ctx context.Context, tx Transaction) ([]Product, error)
	UpdateProductQuantity(ctx context.Context, tx Transaction, productsQuantityMap map[int64]int64) error
}

type Product struct {
	ID          uint
	Name        string
	Description string
	Category    string
	Price       float64
	Is_seasonal bool
	Quantity    int64
}
