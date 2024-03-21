package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).


Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки

*/

var (
	printAfterFlag  int
	printBeforeFlag int
	contextFlag     int
	countFlag       bool
	ignoreCaseFlag  bool
	invertFlag      bool
	fixedFlag       bool
	lineNumFlag     bool
)

func init() {
	flag.IntVar(&printAfterFlag, "A", 0, "lines to print after match")
	flag.IntVar(&printBeforeFlag, "B", 0, "lines to print before match")
	flag.IntVar(&contextFlag, "C", 0, "lines to print around match")
	flag.BoolVar(&countFlag, "c", false, "the number of lines with a match")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "ignore case")
	flag.BoolVar(&invertFlag, "v", false, "exclude")
	flag.BoolVar(&fixedFlag, "F", false, "exact match")
	flag.BoolVar(&lineNumFlag, "n", false, "print the line number")
}

func main() {
	flag.Parse()
	filePath := os.Args[len(os.Args)-1]
	pattern := flag.Args()[:len(flag.Args())-1]
	out := os.Stdout

	if len(os.Args) == 3 {
		err := DefaultGrep(out, filePath, strings.Join(pattern, " "))
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		err := DoGrep(out, filePath, strings.Join(pattern, " "))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func DefaultGrep(out *os.File, filePath, pattern string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	lines := strings.Split(string(file), "\n")

	for _, l := range lines {
		if strings.Contains(l, pattern) {
			_, err = out.WriteString(l + "\n")
			if err != nil {
				fmt.Println("unable to write the found string")
				return err
			}
		}
	}

	return nil
}

func DoGrep(out *os.File, filePath, pattern string) error {
	if printAfterFlag != 0 || printBeforeFlag != 0 || contextFlag != 0 {
		err := printBeforeAfterContext(out, filePath, pattern, printBeforeFlag, printAfterFlag, contextFlag)
		if err != nil {
			fmt.Println("unable to print the necessary number of lines")
			return err
		}
	}
	if countFlag {
		err := printCount(out, filePath, pattern)
		if err != nil {
			fmt.Println("unable to print the line count")
			return err
		}
	}
	if ignoreCaseFlag {
		err := printIgnoreCase(out, filePath, pattern)
		if err != nil {
			fmt.Println("unable to print ignore case")
			return err
		}
	}
	if invertFlag {
		err := printInvert(out, filePath, pattern)
		if err != nil {
			fmt.Println("unable to print inverse")
			return err
		}
	}
	if fixedFlag {
		err := printFixed(out, filePath, pattern)
		if err != nil {
			fmt.Println("unable to print fixed")
			return err
		}
	}
	if lineNumFlag {
		err := printLineNum(out, filePath, pattern)
		if err != nil {
			fmt.Println("unable to print the line number")
			return err
		}
	}

	return nil
}

func printBeforeAfterContext(out *os.File, filePath, pattern string, before, after, context int) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	lines := strings.Split(string(file), "\n")

	if context != 0 {
		after = context
		before = context
	}

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], pattern) {
			for j := before; j > 0; j-- {
				_, err = out.WriteString(lines[i-j] + "\n")
				if err != nil {
					fmt.Println("unable to write the string after")
				}
			}

			_, err = out.WriteString(lines[i] + "\n")
			if err != nil {
				fmt.Println("unable to write the found string")
			}

			if after != 0 {
				for j := 1; j <= after; j++ {
					_, err = out.WriteString(lines[i+j] + "\n")
					if err != nil {
						fmt.Println("unable to write the string after")
					}
				}
			}
		}
	}

	return nil
}

func printCount(out *os.File, filePath, pattern string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	count := 0

	lines := strings.Split(string(file), "\n")

	for i := 0; i < len(lines); i++ {
		if strings.Contains(lines[i], pattern) {
			count++
		}
	}

	_, err = out.WriteString(strconv.Itoa(count))
	if err != nil {
		fmt.Println("unable to write the found string")
	}

	return nil
}

func printIgnoreCase(out *os.File, filePath, pattern string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	lines := strings.Split(string(file), "\n")

	for _, l := range lines {
		if strings.Contains(strings.ToLower(l), pattern) {
			_, err = out.WriteString(l + "\n")
			if err != nil {
				fmt.Println("unable to write the found string")
				return err
			}
		}
	}

	return nil
}

func printInvert(out *os.File, filePath, pattern string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	lines := strings.Split(string(file), "\n")
	withPattern := make(map[int]bool)

	for i, line := range lines {
		if strings.Contains(line, pattern) {
			withPattern[i] = true
		}
	}

	for i, line := range lines {
		if !withPattern[i] {
			_, err = out.WriteString(line + "\n")
			if err != nil {
				fmt.Println("unable to write the found string")
				return err
			}
		}
	}

	return nil
}

func printLineNum(out *os.File, filePath, pattern string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	lines := strings.Split(string(file), "\n")

	for i, l := range lines {
		if strings.Contains(l, pattern) {
			_, err = out.WriteString(strconv.Itoa(i + 1))
			if err != nil {
				fmt.Println("unable to write the number of the found string")
				return err
			}
		}
	}

	return nil
}

func printFixed(out *os.File, filePath, pattern string) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("unable to read the file")
		return nil
	}

	lines := strings.Split(string(file), "\n")

	for i := 0; i < len(lines); i++ {
		if lines[i] == pattern {
			_, err = out.WriteString(lines[i] + "\n")
			if err != nil {
				fmt.Println("unable to write the found string")
				return err
			}
		}
	}

	return nil
}
