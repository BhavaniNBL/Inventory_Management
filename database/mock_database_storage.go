package database

import (
	"fmt"
	"inventory-management/models"
	"time"
)

type MockDatabaseStorage struct {
	Store map[int]models.Product
}

func (db *MockDatabaseStorage) Save(p models.Product) error {
	time.Sleep(1 * time.Second)
	db.Store[p.ID] = p
	return nil
}

func (db *MockDatabaseStorage) GetByID(id int) (*models.Product, error) {
	time.Sleep(1 * time.Second)
	if product, exists := db.Store[id]; exists {
		return &product, nil
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

func (db *MockDatabaseStorage) Delete(id int) error {
	time.Sleep(1 * time.Second)
	if _, exists := db.Store[id]; !exists {
		return fmt.Errorf("product with ID %d not found", id)
	}
	delete(db.Store, id)
	return nil
}

func (db *MockDatabaseStorage) Update(p models.Product) error {
	if _, exists := db.Store[p.ID]; !exists {
		return fmt.Errorf("product with ID %d not found", p.ID)
	}
	db.Store[p.ID] = p
	return nil
}
