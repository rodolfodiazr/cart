package models

type CartBuilder struct {
	cart *Cart
}

func NewCartBuilder() *CartBuilder {
	return &CartBuilder{
		cart: &Cart{},
	}
}

func (b *CartBuilder) AddItem(product Product, quantity int) *CartBuilder {
	b.cart.AddItem(product, quantity)
	return b
}

func (b *CartBuilder) Build() *Cart {
	return b.cart
}
