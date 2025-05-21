package models

import (
	"math"
	"testing"
)

func Test_Cart_AddItem(t *testing.T) {
	tCases := []struct {
		name string

		initialCart   Cart
		itemToAdd     CartItem
		expectedItems []CartItem
	}{
		{
			name: "add a new product to empty cart",

			initialCart: Cart{},
			itemToAdd:   CartItem{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 5},
			expectedItems: []CartItem{
				{Product: Product{ID: "1234", Name: "Product 1", Price: 100.99}, Quantity: 5},
			},
		},
		{
			name: "add a new product to cart with items",

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
			name: "add an existing product to cart",

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
		t.Run(tc.name, func(t *testing.T) {
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

func Test_Cart_Total(t *testing.T) {
	tCases := []struct {
		name string

		cart     Cart
		expected float64
	}{
		{
			name: "cart with no products returns 0.0",

			cart:     Cart{},
			expected: 0.0,
		},
		{
			name: "cart with 1 product: 2x10.4 returns 20.8",

			cart: Cart{
				Items: []CartItem{
					{Product: Product{Name: "Product 1", Price: 10.4}, Quantity: 2},
				},
			},
			expected: 20.8,
		},
		{
			name: "cart with 2 products: 2x10.4 and 4x8.2 returns 53.6",

			cart: Cart{
				Items: []CartItem{
					{Product: Product{Name: "Product 1", Price: 10.4}, Quantity: 2},
					{Product: Product{Name: "Product 2", Price: 8.2}, Quantity: 4},
				},
			},
			expected: 53.6,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			if math.Abs(tc.cart.Total()-tc.expected) >= 0.0001 {
				t.Errorf("expected total %.2f, got %.2f", tc.expected, tc.cart.Total())
			}
		})
	}
}
