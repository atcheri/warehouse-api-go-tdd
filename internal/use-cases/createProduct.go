package usecases

import (
	"fmt"

	"github.com/atcheri/warehouse-api-go-tdd/internal/domain"
	"github.com/google/uuid"
)

type ProductStorer interface {
	Add(domain.Product) error
	FindById(uuid.UUID) (domain.Product, error)
}

type CreateProductUsecase interface {
	Execute(string, float64) (domain.Product, error)
}

type createProduct struct {
	store ProductStorer
}

func NewCreateProductUsecase(store ProductStorer) *createProduct {
	return &createProduct{
		store: store,
	}
}

func (uc createProduct) Execute(name string, prince float64) (domain.Product, error) {
	product := domain.Product{
		ID:    uuid.New(),
		Name:  name,
		Price: prince,
	}

	err := uc.store.Add(product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error executing the create product usecase: %w", err)
	}

	return product, nil
}
