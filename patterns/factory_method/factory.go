package main

import "fmt"

// getCar возвращает нужную структуру в зависимости от входящего параметра - типа машины
func getCar(carType string) (CarI, error) {
	if carType == "Suzuki" {
		return newSuzuki(), nil
	}
	if carType == "Opel" {
		return newOpel(), nil
	}
	return nil, fmt.Errorf("wrong car type passed")
}
