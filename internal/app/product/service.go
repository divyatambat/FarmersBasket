package product

import (
	"context"

	"github.com/divyatambat/FarmersBasket/internal/pkg/apperrors"
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"

	"github.com/divyatambat/FarmersBasket/internal/repository"
)

type service struct {
	productRepo repository.ProductStorer
}

type Service interface {
	GetProductByID(ctx context.Context, productID int64) (dto.Product, error)
	ListProducts(ctx context.Context) ([]dto.Product, error)
	UpdateProductQuantity(ctx context.Context, productsQuantityMap map[int64]int64) error
}

func NewService(productRepo repository.ProductStorer) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (ps *service) GetProductByID(ctx context.Context, productID int64) (dto.Product, error) {
	productInfoDB, err := ps.productRepo.GetProductByID(ctx, nil, productID)
	if err != nil {
		return dto.Product{}, nil
	}

	if productInfoDB.ID == 0 {
		return dto.Product{}, apperrors.ProductNotFound{ID: productID}
	}

	productInfo := MapRepoObjectToDto(productInfoDB)
	return productInfo, nil
}

func (ps *service) ListProducts(ctx context.Context) ([]dto.Product, error) {
	products := make([]dto.Product, 0)

	productsListDB, err := ps.productRepo.ListProducts(ctx, nil)
	if err != nil {
		return products, err
	}

	for _, productInfo := range productsListDB {
		products = append(products, MapRepoObjectToDto(productInfo))
	}

	return products, nil
}

type UpdateProductQuantity struct {
	ProductID int64
	Quantity  int64
}

func (ps *service) UpdateProductQuantity(ctx context.Context, productsQuantityMap map[int64]int64) error {
	err := ps.productRepo.UpdateProductQuantity(ctx, nil, productsQuantityMap)
	return err
}
