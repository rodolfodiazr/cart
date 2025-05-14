package models

type Cart struct {
	Items []CartItem
}

// AddItem adds a product to the cart
func (c *Cart) AddItem(product Product, quantity int) {
	for i, item := range c.Items {
		if item.Product.ID == product.ID {
			c.Items[i].Quantity += quantity
			return
		}
	}

	c.Items = append(c.Items, CartItem{Product: product, Quantity: quantity})
}

// Total calculates the total of all products in the cart
func (c *Cart) Total() float64 {
	var total float64
	for _, item := range c.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}
