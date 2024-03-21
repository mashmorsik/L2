package main

// Car - структура, которая имплементирует интерфейс CarI
type Car struct {
	model   string
	country string
}

func (c *Car) setModel(model string) {
	c.model = model
}

func (c *Car) getModel() string {
	return c.model
}

func (c *Car) setCountry(country string) {
	c.country = country
}

func (c *Car) getCountry() string {
	return c.country
}
