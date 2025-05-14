package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rodolfodiazr/cart/models"
)

func main() {
	product := models.Product{
		ID:    uuid.Must(uuid.NewV6()).String(),
		Name:  "Keyboard",
		Price: 20,
	}

	cart := models.NewCartBuilder().
		AddItem(product, 1).
		AddItem(product, 2).
		Build()

	fmt.Print("CART:\n\n")
	for i, item := range cart.Items {
		fmt.Printf("%d  %v  %d  $%.1f  $%.1f\n", i+1, item.Product.Name, item.Quantity, item.Product.Price, (float64(item.Quantity) * item.Product.Price))
	}
	fmt.Printf("\nTOTAL: $%.1f\n", cart.Total())
}
