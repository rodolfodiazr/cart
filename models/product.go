package models

// Product represents an item available for purchase.
// It includes a unique identifier, a name, and a price.
type Product struct {
	ID    string
	Name  string
	Price float64
}
