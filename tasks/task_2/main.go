package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
}

func UnpackString(str string) (string, error) {
	runes := []rune(str)
	var sb strings.Builder

	if str == "" {
		return "", errors.New("invalid string")
	}

	for i := 0; i < len(runes)-1; i++ {
		first := runes[i]
		second := runes[i+1]
		switch {
		case !unicode.IsDigit(first) && unicode.IsDigit(second):
			num, _ := strconv.Atoi(string(second))
			sb.WriteString(strings.Repeat(string(first), num))
		case !unicode.IsDigit(first) && !unicode.IsDigit(second):
			sb.WriteString(string(first))
		}
	}

	last := runes[len(runes)-1]
	if !unicode.IsDigit(last) {
		sb.WriteString(string(last))
	}

	return sb.String(), nil
}
