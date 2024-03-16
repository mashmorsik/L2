package main

//Что выведет программа? Объяснить вывод программы.

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// программа выведет все значения от 0 до 9, а потом будет deadlock, тк основная горутина будет ждать новой записи
	// в канал или закрытия канала, но этого не происходит
	for n := range ch {
		println(n)
	}
}