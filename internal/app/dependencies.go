package app

import (
	"database/sql"

	"github.com/divyatambat/FarmersBasket/internal/app/order"
	"github.com/divyatambat/FarmersBasket/internal/app/product"
	"github.com/divyatambat/FarmersBasket/internal/app/user"
	repository "github.com/divyatambat/FarmersBasket/internal/repository/postgres"
)

type Dependencies struct {
	OrderService   order.Service
	ProductService product.Service
	UserService    user.Service
}

func NewServices(db *sql.DB) (*Dependencies, error) {
	// Initialize database-specific repositories
	orderRepo := repository.NewOrderRepo(db)
	orderItemsRepo := repository.NewOrderItemRepo(db)
	productRepo := repository.NewProductRepo(db)
	userRepo := repository.NewUserRepo(db)

	// Initialize services with repository dependencies
	productService := product.NewService(productRepo)
	orderService := order.NewService(orderRepo, orderItemsRepo, productService)
	userService := user.NewService(userRepo)

	return &Dependencies{
		OrderService:   orderService,
		ProductService: productService,
		UserService:    userService,
	}, nil
}
