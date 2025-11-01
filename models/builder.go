package models

// CartBuilder provides a convenient way to construct a Cart instance
// using a fluent interface. It allows adding products step by step
// before building the final Cart.
type CartBuilder struct {
	cart *Cart
}

// NewCartBuilder creates and returns a new CartBuilder instance
// with an empty Cart ready for configuration.
func NewCartBuilder() *CartBuilder {
	return &CartBuilder{
		cart: &Cart{},
	}
}

// AddItem adds a product to the cart with the specified quantity.
// If the quantity is zero or negative, the item is ignored.
// Returns the same CartBuilder to allow method chaining.
func (b *CartBuilder) AddItem(product Product, quantity int) *CartBuilder {
	if quantity <= 0 {
		return b
	}

	b.cart.AddItem(product, quantity)
	return b
}

// Build returns the fully constructed Cart instance.
func (b *CartBuilder) Build() *Cart {
	return b.cart
}
