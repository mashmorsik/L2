package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Что выведет программа? Объяснить вывод программы.

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()

	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)

	// Программа выведет числа, переданные в каналы a и b (в случайном порядке, будет зависеть от времени, которое
	// горутины будут тратить на sleep), далее программа будет бесконечно выводить 0, так как такое значение мы
	// получаем при чтении из закрытого канала
	for v := range c {
		fmt.Println(v)
	}
}
