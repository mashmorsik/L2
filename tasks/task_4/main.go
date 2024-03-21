package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
Поиск анаграмм по словарю
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству, 'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

*/

var list = []string{"горбик", "Гробик", "керамит", "собака", "Метрика", "Материк", "Грибок"}

func main() {
	fmt.Println(AnagramDict(list))
}

func AnagramDict(list []string) map[string][]string {
	var (
		dict, tempMap, added = make(map[string][]string), make(map[int][]int), make(map[int]bool)
		sortedCharArr        []string
	)

	for i := 0; i < len(list); i++ {
		charArr := strings.Split(strings.ToLower(list[i]), "")
		slices.Sort(charArr)
		sortedCharArr = append(sortedCharArr, strings.Join(charArr, ""))
	}

	for ind, v := range sortedCharArr {
		for i := 0; i < len(sortedCharArr); i++ {
			if v == sortedCharArr[i] && !added[i] {
				tempMap[ind] = append(tempMap[ind], i)
				added[i] = true
			}
		}
	}

	for _, v := range tempMap {
		if len(v) > 1 {
			for _, w := range v[1:] {
				dict[strings.ToLower(list[v[0]])] = append(dict[strings.ToLower(list[v[0]])], strings.ToLower(list[w]))
			}
		}
	}

	for _, v := range dict {
		slices.Sort(v)
	}
	fmt.Println(sortedCharArr)
	return dict
}
