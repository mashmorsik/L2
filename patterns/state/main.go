package main

import "fmt"

/*
Паттерн Состояние

Состояние - поведенческий паттерн проектирования, позволяющий объекту
изменять своё поведение в зависимости от внутреннего состояния. При этом
объект будет выглядеть так, будто изменился его класс.

Плюсы:
- Избавляет от множества больших условных операторов машины состояний;
- Концентрирует в одном месте код, связанный с определённым состоянием;
- Упрощает код контекстаю

Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.

Пример использования: пользовательские сессии (пользователь незалогинился, залогинился, разлогинился).
*/

func main() {
	sessionManager := NewSessionManager()

	err := sessionManager.Authenticate("user123", "password123")
	if err != nil {
		fmt.Println("Authentication failed:", err)
		return
	}

	err = sessionManager.AccessResource("example/resource")
	if err != nil {
		fmt.Println("Access denied:", err)
		return
	}

	err = sessionManager.Logout()
	if err != nil {
		fmt.Println("Logout failed:", err)
		return
	}
}
