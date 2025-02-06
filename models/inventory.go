package models

import "fmt"

type Inventory struct {
	Products map[int]Product
}

func (inv *Inventory) AddProduct(p Product) error {
	if _, exists := inv.Products[p.ID]; exists {
		return fmt.Errorf("product with ID %d already exists", p.ID)
	}
	inv.Products[p.ID] = p
	return nil
}

func (inv *Inventory) RemoveProduct(id int) error {
	if _, exists := inv.Products[id]; !exists {
		return fmt.Errorf("product with ID %d not found", id)
	}

	delete(inv.Products, id)
	return nil
}

func (inv *Inventory) FindProductByName(name string) (*Product, error) {
	for _, product := range inv.Products {
		if product.Name == name {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product with name %s not found", name)
}

func (inv *Inventory) FindProductByID(id int) (*Product, error) {
	product, exists := inv.Products[id]
	if !exists {
		return nil, fmt.Errorf("product with ID %d not found", id)
	}
	return &product, nil
}

func (inv *Inventory) ListByCategory(category string) []Product {
	var result []Product
	for _, product := range inv.Products {
		if product.Category == category {
			result = append(result, product)
		}
	}

	return result
}

func (inv *Inventory) TotalValue() float64 {
	total := 0.0
	for _, product := range inv.Products {
		total += product.Price * float64(product.Quantity)
	}
	return total
}

func (inv *Inventory) GetAllProducts() []Product {
	products := []Product{}
	for _, product := range inv.Products {
		products = append(products, product)
	}
	return products

}

func (inv *Inventory) UpdateProduct(p Product) error {
	if _, exists := inv.Products[p.ID]; !exists {
		return fmt.Errorf("product with ID %d not found", p.ID)
	}
	inv.Products[p.ID] = p
	return nil
}
