package main

import "fmt"

/*
Паттерн Строитель

Паттерн «строитель» является порождающим паттерном проектирования,
предназначен для создания сложных объектов, многосоставных объектов.

Плюсы:
- Позволяет создавать продукты пошагово;
- Позволяет использовать один и тот же код для создания различных продуктов;
- Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:
- Усложняет код программы из-за введения дополнительных структур;
- Клиент будет привязан к конкретным структурам строителей,
так как в интерфейсе директора может не быть метода получения результата.

Пример использования паттерна: создание структуры Product.
*/

func main() {
	builder1 := NewProductBuilder()
	headphones := builder1.SetName("Headphones").SetDescription("New extra awesome.").
		SetQuantity("Khimki_i1", 32).SetQuantity("Moscow_p7", 12).BuildProduct()

	builder2 := NewProductBuilder()
	table := builder2.SetName("Table").SetDescription("Comfy wooden large.").BuildProduct()

	fmt.Println(headphones)
	fmt.Println(table)
}
