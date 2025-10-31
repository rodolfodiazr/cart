package models

// CartItem represents a product and its quantity within a shopping cart.
type CartItem struct {
	Product  Product
	Quantity int
}
