package main

// ErrorHandler интерфейс представляет обработчик ошибок
type ErrorHandler interface {
	SetNext(handler ErrorHandler) // Установить следующий обработчик в цепочке
	Handle(err error)             // Обработать ошибку
}
