package main

type Opel struct {
	Car
}

// newOpel создает экземпляр структуры Opel, которая "наследует" структуру Car с ее методами
func newOpel() CarI {
	return &Opel{
		Car: Car{
			model:   "Opel",
			country: "Germany",
		},
	}
}
