package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/divyatambat/FarmersBasket/internal/repository"
)

type productStore struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) repository.ProductStorer {
	return &productStore{db: db}
}

func (ps *productStore) GetProductByID(ctx context.Context, tx repository.Transaction, productID int64) (repository.Product, error) {
	var product repository.Product

	row := ps.db.QueryRowContext(ctx, `SELECT * FROM product WHERE id = $1`, productID)
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Category, &product.Price, &product.Is_seasonal, &product.Quantity)
	if err == sql.ErrNoRows {
		return repository.Product{}, errors.New("product not found")
	} else if err != nil {
		return repository.Product{}, fmt.Errorf("error getting product: %w", err)
	}

	return product, nil
}

func (ps *productStore) ListProducts(ctx context.Context, tx repository.Transaction) ([]repository.Product, error) {
	productList := make([]repository.Product, 0)

	rows, err := ps.db.QueryContext(ctx, `SELECT * FROM product`)
	if err != nil {
		return productList, fmt.Errorf("error listing products: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product repository.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Category, &product.Price, &product.Is_seasonal, &product.Quantity)
		if err != nil {
			return productList, fmt.Errorf("error iterating product rows: %w", err)
		}
		productList = append(productList, product)
	}

	return productList, nil
}

func (ps *productStore) UpdateProductQuantity(ctx context.Context, tx repository.Transaction, productsQuantityMap map[int64]int64) error {
	// Check if transaction is available for prepared statements
	txStmt, ok := tx.(interface {
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	})
	if !ok {
		return errors.New("transaction does not support prepared statements")
	}

	for productID, quantity := range productsQuantityMap {
		_, err := txStmt.ExecContext(ctx, `UPDATE products SET quantity = $1 WHERE id = 3`, quantity, productID)
		if err != nil {
			return fmt.Errorf("error updating product quantity: %w", err)
		}
	}

	return nil
}
