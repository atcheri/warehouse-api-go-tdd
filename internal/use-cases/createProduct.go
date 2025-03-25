package usecases

import (
	"github.com/atcheri/warehouse-api-go-tdd/internal/domain"
	"github.com/google/uuid"
)

type ProductCreator interface {
}

type CreateProduct struct {
	// store ProductCreator
}

func (uc CreateProduct) Execute(name string, prince float64) domain.Product {
	return domain.Product{
		ID:    uuid.New(),
		Name:  name,
		Price: prince,
	}
}
