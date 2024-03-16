package main

import (
	"fmt"
	"strings"
)

// Visitor определяет интерфейс посетителя
type Visitor interface {
	Visit(data Data) error
}

// Data определяет данные для валидации
type Data struct {
	Username string
	Email    string
	Password string
}

// ValidatorVisitor реализует посетителя для валидации данных
type ValidatorVisitor struct{}

// Visit выполняет валидацию данных
func (v *ValidatorVisitor) Visit(data Data) error {
	if data.Username == "" {
		return fmt.Errorf("username is required")
	}
	if len(data.Password) < 6 {
		return fmt.Errorf("password should be at least 6 characters long")
	}
	if !strings.Contains(data.Email, "@") {
		return fmt.Errorf("emails must contain @")
	}

	return nil
}

// Validate функция для запуска валидации данных
func Validate(data Data, visitor Visitor) error {
	return visitor.Visit(data)
}
