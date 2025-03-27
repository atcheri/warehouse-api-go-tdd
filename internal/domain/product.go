package domain

import "github.com/google/uuid"

type Product struct {
	ID    uuid.UUID
	Name  string
	Price float64
}

func NewProduct(name string, price float64) Product {
	return Product{
		ID:    uuid.New(),
		Name:  name,
		Price: price,
	}
}
