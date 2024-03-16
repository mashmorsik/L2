package main

import "fmt"

// Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1

	// выводит: 2, тк в defer мы инкрементируем заранее проинициализированную переменную (x int)
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1

	// выводит: 1, тк переменная x объявляется внутри функции и возвращается в return, а ее инкремент
	// происходит уже в defer, поэтому его результат не виден в выводе
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
