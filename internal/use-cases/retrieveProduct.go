package usecases

import (
	"fmt"

	"github.com/atcheri/warehouse-api-go-tdd/internal/domain"
	"github.com/google/uuid"
)

type RetrieveProductUsecase interface {
	Execute(uuid.UUID) (domain.Product, error)
}

type retriveProduct struct {
	store ProductStorer
}

func NewRetrieveProductUsecase(store ProductStorer) *retriveProduct {
	return &retriveProduct{
		store: store,
	}
}

func (uc retriveProduct) Execute(id uuid.UUID) (domain.Product, error) {
	product, err := uc.store.FindById(id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error executing the retrive product usecase: %w", err)
	}

	return product, nil
}
