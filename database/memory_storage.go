package database

import (
	"fmt"
	"inventory-management/models"
)

type MemoryStorage struct {
	Store map[int]models.Product
}

func (ms *MemoryStorage) Save(p models.Product) error {
	ms.Store[p.ID] = p
	return nil
}

func (ms *MemoryStorage) GetByID(id int) (*models.Product, error) {
	if product, exists := ms.Store[id]; exists {
		return &product, nil
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

func (ms *MemoryStorage) ListByCategory(category string) []models.Product {
	var filteredProducts []models.Product
	for _, product := range ms.Store {
		if product.Category == category {
			filteredProducts = append(filteredProducts, product)
		}
	}

	return filteredProducts

}

func (ms *MemoryStorage) Delete(id int) error {
	if _, exists := ms.Store[id]; !exists {
		return fmt.Errorf("product with ID %d not found", id)
	}
	delete(ms.Store, id)
	return nil
}

func (ms *MemoryStorage) Update(p models.Product) error {
	if _, exists := ms.Store[p.ID]; !exists {
		return fmt.Errorf("product with ID %d not found", p.ID)
	}
	ms.Store[p.ID] = p
	return nil
}
