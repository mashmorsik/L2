package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
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

var (
	columnFlag  int
	numericFlag bool
	reverseFlag bool
	uniqueFlag  bool
)

func init() {
	flag.IntVar(&columnFlag, "k", 0, "column for sorting")
	flag.BoolVar(&numericFlag, "n", false, "sort by number")
	flag.BoolVar(&reverseFlag, "r", false, "sort in reverse order")
	flag.BoolVar(&uniqueFlag, "u", false, "show only unique values")
}

func main() {
	flag.Parse()
	filePath := os.Args[len(os.Args)-1]
	out := os.Stdout

	if len(os.Args) == 2 {

		err := DefaultSort(out, filePath)
		if err != nil {
			fmt.Println("unable to sort by default")
			return
		}
	} else {
		err := DoSort(out, filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func DoSort(out *os.File, filePath string) error {
	if columnFlag != 0 {
		err := ColumnsSort(out, filePath, columnFlag)
		if err != nil {
			fmt.Println("unable to sort by default")
			return err
		}
	}
	if numericFlag {
		err := NumericSort(out, filePath)
		if err != nil {
			fmt.Println("unable to do numeric sort")
			return err
		}
	}
	if reverseFlag {
		err := ReverseSort(out, filePath)
		if err != nil {
			fmt.Println("unable to do reverse sort")
			return err
		}
	}
	if uniqueFlag {
		err := UniqueSort(out, filePath)
		if err != nil {
			fmt.Println("unable to do unique sort")
			return err
		}
	}

	return nil
}

func ColumnsSort(out *os.File, filePath string, column int) error {
	lines, err := ReadLines(filePath, "\n")

	columnInd := column - 1
	if err != nil {
		fmt.Println("unable to convert column to int")
		return err
	}

	sort.Slice(lines, func(i, j int) bool {
		columnValues1 := strings.Fields(lines[i])
		columnValues2 := strings.Fields(lines[j])

		if len(columnValues1) <= columnInd || len(columnValues2) <= columnInd {
			return false
		}

		value1 := columnValues1[columnInd]
		value2 := columnValues2[columnInd]

		return value1 < value2
	})

	err = WriteLines(out, lines)
	if err != nil {
		fmt.Println("unable to write output")
		return err
	}

	return nil
}

func NumericSort(out *os.File, filePath string) error {
	lines, err := ReadLines(filePath, " ")
	var intArr []int

	for _, l := range lines {
		lint, err := strconv.Atoi(l)
		if err != nil {
			return errors.New("file needs to contain integers")
		}
		intArr = append(intArr, lint)
	}
	slices.Sort(lines)

	err = WriteLines(out, lines)
	if err != nil {
		fmt.Println("unable to write output")
		return err
	}

	return nil
}

func ReverseSort(out *os.File, filePath string) error {
	lines, err := ReadLines(filePath, " ")
	slices.Sort(lines)
	slices.Reverse(lines)

	err = WriteLines(out, lines)
	if err != nil {
		fmt.Println("unable to write output")
		return err
	}

	return nil
}

func UniqueSort(out *os.File, filePath string) error {
	lines, err := ReadLines(filePath, " ")
	elements := make(map[string]struct{})
	var sb strings.Builder

	for _, e := range lines {
		_, ok := elements[e]
		if !ok {
			elements[e] = struct{}{}
			sb.WriteString(e + " ")
		}
	}

	err = WriteLines(out, strings.Split(sb.String(), " "))
	if err != nil {
		fmt.Println("unable to write output")
		return err
	}

	return nil
}

func DefaultSort(out *os.File, filePath string) error {
	lines, err := ReadLines(filePath, " ")
	slices.Sort(lines)

	err = WriteLines(out, lines)
	if err != nil {
		fmt.Println("unable to write output")
		return err
	}

	return nil
}

// additional functions to read and write lines
func ReadLines(filePath, split string) ([]string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to open the file")
		return nil, err
	}

	lines := strings.Split(string(file), split)
	return lines, nil
}

func WriteLines(out *os.File, lines []string) error {
	_, err := out.WriteString(strings.Join(lines, " "))
	if err != nil {
		fmt.Println("unable to write the sorted string")
		return err
	}
	return nil
}
