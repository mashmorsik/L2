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
	fmt.Println(err)
	fmt.Println(err == nil)
}
