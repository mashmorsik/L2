package main

import (
	"fmt"
	"time"
)

/*
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих
каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь,
однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме. В этом случае удобнее
использовать вызов единственной функции, которая, приняв на вход один или более or-каналов,
реализовывала бы весь функционал.
*/

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

// or принимает переменное количество каналов для чтения и возвращает канал,
// который закрывается, как только один из переданных каналов закрывается.
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если не передано ни одного канала, возвращается nil.
		return nil
	case 1:
		// Если передан только один канал, он возвращается как результат.
		return channels[0]
	}

	// Создаем канал, который будет возвращен как результат функции.
	chDone := make(chan interface{})

	// Запускаем горутину для отслеживания закрытия каналов и закрытия chDone.
	go func() {
		defer close(chDone)

		switch len(channels) {
		case 2:
			// Если передано два канала, блокируемся до закрытия одного из них.
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			// Если передано более двух каналов, используем рекурсию для вызова or с оставшимися каналами.
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], chDone)...):
			}
		}
	}()

	return chDone
}
