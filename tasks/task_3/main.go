package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

*/

func main() {
	filePath := os.Args[1]

	err := Sort(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Sort(filePath string) error {
	out := os.Stdout

	if len(os.Args) == 2 {
		err := DefaultSort(out, filePath)
		if err != nil {
			fmt.Println("unable to sort by default")
			return err
		}
	}

	return nil
}

func DefaultSort(out *os.File, filePath string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to open the file")
		return err
	}

	arr := strings.Split(string(file), "")
	slices.Sort(arr)

	_, err = out.WriteString(strings.Join(arr, "\n"))
	if err != nil {
		fmt.Println("unable to write the sorted string")
		return err
	}
	return nil
}
