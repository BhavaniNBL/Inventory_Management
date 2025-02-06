package api

import (
	"fmt"
	"inventory-management/models"
	"math/rand"
	"time"
)

type ExternalAPI struct {
	BaseURL string
}

func (extAPI *ExternalAPI) FetchProductDetails(id int) (*models.Product, error) {

	time.Sleep(time.Duration(rand.Intn(4)) * time.Second)

	if rand.Intn(10) == 0 {
		return nil, fmt.Errorf("failed to fetch product details from external API")
	}

	product := &models.Product{
		ID:       id,
		Name:     fmt.Sprintf("Product %d", id),
		Price:    float64(rand.Intn(1000) + 100),
		Quantity: rand.Intn(100),
		Category: "Electronics",
	}

	return product, nil
}
