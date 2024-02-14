package product

import (
	"github.com/divyatambat/FarmersBasket/internal/pkg/dto"
	repository "github.com/divyatambat/FarmersBasket/internal/repository"
)

type ProductType string

const (
	vegetable ProductType = "Vegetables"
	fruits    ProductType = "Fruits"
)

func MapRepoObjectToDto(repoObj repository.Product) dto.Product {
	return dto.Product{
		ID:          int64(repoObj.ID),
		Name:        repoObj.Name,
		Description: repoObj.Description,
		Category:    repoObj.Category,
		Price:       repoObj.Price,
		Is_seasonal: repoObj.Is_seasonal,
		Quantity:    int64(repoObj.Quantity),
	}
}

func MapDtoObjectToRepo(product dto.Product) repository.Product {
	return repository.Product{
		Name:        product.Name,
		Description: product.Category,
		Price:       product.Price,
		Category:    product.Category,
		Quantity:    product.Quantity,
	}
}
