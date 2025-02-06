package models

import "fmt"

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"category"`
}

func (p *Product) UpdatePrice(newPrice float64) {
	p.Price = newPrice
}

func (p *Product) Sell(quantity int) error {
	if p.Quantity < quantity {
		return fmt.Errorf("insufficient stock: only %d items available", p.Quantity)
	}

	p.Quantity -= quantity
	return nil
}

func (p *Product) Restock(quantity int) {
	p.Quantity += quantity
}

func (p *Product) Display() string {
	return fmt.Sprintf("ID: %d, Name: %s, Price: %.2f, Quantity: %d, Category: %s", p.ID, p.Name, p.Price, p.Quantity, p.Category)
}
