package main

import "fmt"

// Что выведет программа? Объяснить вывод программы.

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	// []int можно опустить
	var b []int = a[1:4]

	// программа выведет [77, 78, 79], тк мы берем слайс от массива начиная с 1 индекса и до 4 не включительно
	fmt.Println(b)
}
