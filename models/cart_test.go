package models

import (
	"testing"
)

func Test_Cart_AddItem(t *testing.T) {
	tCases := []struct {
		description string

		initialCart   Cart
		itemToAdd     CartItem
		expectedItems []CartItem
	}{
		{
			description: "Add a new product to empty cart.",

			initialCart: Cart{},
			itemToAdd:   CartItem{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 5},
			expectedItems: []CartItem{
				{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 5},
			},
		},
		{
			description: "Add a new product to cart.",

			initialCart: Cart{
				Items: []CartItem{
					{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 4},
				},
			},
			itemToAdd: CartItem{Product: Product{ID: "5678", Name: "Product 5", Price: 80.5}, Quantity: 2},
			expectedItems: []CartItem{
				{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 4},
				{Product: Product{ID: "5678", Name: "Product 5", Price: 80.5}, Quantity: 2},
			},
		},
		{
			description: "Add an existing product to cart.",

			initialCart: Cart{
				Items: []CartItem{
					{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 2},
				},
			},
			itemToAdd: CartItem{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 8},
			expectedItems: []CartItem{
				{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 10},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.description, func(t *testing.T) {
			cart := tc.initialCart
			cart.AddItem(tc.itemToAdd.Product, tc.itemToAdd.Quantity)

			if len(cart.Items) != len(tc.expectedItems) {
				t.Fatalf("expected %d items, got %d", len(tc.expectedItems), len(cart.Items))
			}

			actualItems := make(map[string]CartItem)
			for _, item := range cart.Items {
				actualItems[item.Product.ID] = item
			}

			for _, expected := range tc.expectedItems {
				got, ok := actualItems[expected.Product.ID]
				if !ok {
					t.Errorf("expected product ID %s not found", expected.Product.ID)
					continue
				}

				if got.Quantity != expected.Quantity {
					t.Errorf("product ID %s: expected quantity %d, got %d", expected.Product.ID, expected.Quantity, got.Quantity)
				}
			}
		})
	}
}
