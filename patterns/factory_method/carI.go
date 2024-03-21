package main

// CarI - интерфейс с методами для объектов фабрики
type CarI interface {
	setModel(name string)
	setCountry(power string)
	getModel() string
	getCountry() string
}
