package db

import (
	"errors"
	"fmt"
	"sync"

	"github.com/atcheri/warehouse-api-go-tdd/internal/domain"
	"github.com/google/uuid"
)

type inMemoryDB struct {
	products map[string]domain.Product
	mutex    sync.Mutex
}

func NewInMemoryDB() *inMemoryDB {
	return &inMemoryDB{
		products: make(map[string]domain.Product),
	}
}

// Add adds a product into the product store.
func (i *inMemoryDB) Add(p domain.Product) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, ok := i.products[p.Name]; ok {
		return errors.New("product already exists")
	}

	i.products[p.Name] = p

	return nil
}

func (i *inMemoryDB) FindById(id uuid.UUID) (domain.Product, error) {
	for _, p := range i.products {
		if p.ID == id {
			return p, nil
		}
	}

	return domain.NewProduct("not-found", 0), fmt.Errorf("product not found with id: %s", id.String())
}
