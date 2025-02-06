package main

import (
	"fmt"
	"inventory-management/database"
	"inventory-management/models"
	"inventory-management/routes"
	"log"
	"net/http"
)

func main() {
	inv := &models.Inventory{Products: make(map[int]models.Product)}
	storage := &database.MemoryStorage{Store: make(map[int]models.Product)}

	routes.InitializeRoutes(inv, *storage)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
