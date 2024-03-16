package main

import (
	"image"
)

// Product - финальная структура продукта
type Product struct {
	Name        string
	Description string
	Photo       image.Image
	Quantity    map[string]int
}

// ProductBuilderI - интерфейс для создания продукта
type ProductBuilderI interface {
	SetName(name string) ProductBuilderI
	SetDescription(description string) ProductBuilderI
	SetPhoto(photo image.Image) ProductBuilderI
	SetQuantity(stockId string, quantity int) ProductBuilderI
	BuildProduct() Product
}

// ProductBuilder - структура продукта для строительства
type ProductBuilder struct {
	Name        string
	Description string
	Photo       image.Image
	Quantity    map[string]int
}

func NewProductBuilder() ProductBuilderI {
	return &ProductBuilder{
		Quantity: make(map[string]int),
	}
}

// SetName записывает имя продукта
func (p *ProductBuilder) SetName(name string) ProductBuilderI {
	p.Name = name
	return p
}

// SetDescription записывает описание продукта
func (p *ProductBuilder) SetDescription(description string) ProductBuilderI {
	p.Description = description
	return p
}

// SetPhoto прикрепляет фотографию продукта
func (p *ProductBuilder) SetPhoto(photo image.Image) ProductBuilderI {
	p.Photo = photo
	return p
}

// SetQuantity указывает количество продукта на конкретных складах
func (p *ProductBuilder) SetQuantity(stockId string, quantity int) ProductBuilderI {
	p.Quantity[stockId] = quantity
	return p
}

// BuildProduct возвращает продукт с заполненными полями
func (p *ProductBuilder) BuildProduct() Product {
	return Product{
		Name:        p.Name,
		Description: p.Description,
		Photo:       p.Photo,
		Quantity:    p.Quantity,
	}
}
