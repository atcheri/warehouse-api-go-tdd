package db

import (
	"errors"

	"github.com/atcheri/warehouse-api-go-tdd/internal/domain"
)

type inMemoryDB struct {
	products map[string]domain.Product
}

func NewInMemoryDB() *inMemoryDB {
	return &inMemoryDB{
		products: make(map[string]domain.Product),
	}
}

// Add adds a product into the product store.
func (i *inMemoryDB) Add(p domain.Product) error {
	if _, ok := i.products[p.Name]; ok {
		return errors.New("product already exists")
	}

	i.products[p.Name] = p

	return nil
}
