package main

// сделать

import (
	"fmt"
	"os"
)

// Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов
//и их отличие от пустых интерфейсов.

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	// выведет nil
	fmt.Println(err)
	// Поскольку err имеет тип *os.PathError и его значение не nil (т.е. он содержит нулевой указатель),
	// сравнение err == nil возвращает false.
	fmt.Println(err == nil)
}
