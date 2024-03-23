package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/* Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN,
разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

*/

var (
	fieldsFlag    int
	delimiterFlag string
	separatedFlag bool
)

func init() {
	flag.IntVar(&fieldsFlag, "f", 0, "to choose fields/columns")
	flag.StringVar(&delimiterFlag, "d", "", "use another delimiter")
	flag.BoolVar(&separatedFlag, "s", false, "only separated strings")
}

func main() {
	flag.Parse()

	var delimiter string

	if delimiterFlag != "" {
		delimiter = delimiterFlag
	} else {
		delimiter = "\t"
	}

	lines := readLines(delimiter)

	writeResponse(delimiter, lines)
}

func readLines(delimiter string) [][]string {
	var lines [][]string

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if separatedFlag && !strings.Contains(line, delimiter) {
			continue
		}
		lines = append(lines, strings.Split(line, delimiter))
	}

	return lines
}

func writeResponse(delimiter string, lines [][]string) {
	var sb strings.Builder

	if fieldsFlag != 0 {
		for i, line := range lines {
			if i == len(lines)-1 {
				sb.WriteString(line[fieldsFlag-1])
			} else {
				sb.WriteString(line[fieldsFlag-1] + delimiter)
			}
		}
	}
	fmt.Println(sb.String())
	return
}
