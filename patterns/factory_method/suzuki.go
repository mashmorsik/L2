package main

type Suzuki struct {
	Car
}

// newSuzuki создает экземпляр структуры Suzuki, которая "наследует" структуру Car с ее методами
func newSuzuki() CarI {
	return &Suzuki{
		Car: Car{
			model:   "Suzuki",
			country: "Japan",
		},
	}
}
