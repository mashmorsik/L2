package main

import (
	"fmt"
	"time"
)

/*
Паттерн Стратегия

Стратегия - поведенческий паттерн проектирования, выделяющий набор алгоритмов,
решающих конкретную задачу. Позволяет выбирать алгоритм
в процессе выполнения программы.

Плюсы:
- Позволяет изменять поведение объекта во время выполнения программы, не
изменяя его структуры;
- Уменьшает зависимость между клиентом и объектами, предоставляя клиентскому коду
в качестве зависимости абстракцию;
- Позволяет добавлять новые алгоритмы без изменения существующих и выбирать наиболее
подходящий алгоритм из набора во время выполнения.

Минусы:
- Может усложнить код программы за счёт дополнительных классов и
интерфейсов стратегий;
- В случае если стратегии схожи, код может стать избыточным и повторяющимся.
Клиент должен знать о существовании стратегий и уметь выбирать их.

Пример использования паттерна: генерация отчетов в разных форматах (CSV, PDF, HTML).
*/

func main() {
	newReport := &Report{
		Id:          "HGB-567",
		Name:        "Expenses report",
		DateCreated: time.Now(),
		Type:        2,
		Description: "Very detailed description",
		Author:      "Nick Ivanov",
		Status:      1,
	}

	CSV := &CSVReport{}
	context := &ReportContext{Strategy: CSV}
	fmt.Println(context.Execute(newReport))

	PDF := &PDFReport{}
	context = &ReportContext{PDF}
	fmt.Println(context.Execute(newReport))

	HTML := &HTMLReport{}
	context = &ReportContext{Strategy: HTML}
	fmt.Println(context.Execute(newReport))
}
