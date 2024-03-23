package main

import (
	"errors"
	"fmt"
)

// BaseHandler базовый обработчик ошибок
type BaseHandler struct {
	next ErrorHandler
}

func (h *BaseHandler) SetNext(handler ErrorHandler) {
	h.next = handler
}

func (h *BaseHandler) Handle(err error) {
	if h.next != nil {
		h.next.Handle(err)
	} else {
		fmt.Println("No handler can process the error")
	}
}

// AuthenticationErrorHandler обработчик ошибок аутентификации
type AuthenticationErrorHandler struct {
	BaseHandler
}

func (h *AuthenticationErrorHandler) Handle(err error) {
	if errors.Is(err, ErrAuthenticationFailed) {
		fmt.Println("Authentication error:", err)
	} else {
		h.BaseHandler.Handle(err)
	}
}

// DatabaseErrorHandler обработчик ошибок базы данных
type DatabaseErrorHandler struct {
	BaseHandler
}

func (h *DatabaseErrorHandler) Handle(err error) {
	if errors.Is(err, ErrDatabaseConnectionFailed) {
		fmt.Println("Database connection error:", err)
	} else {
		h.BaseHandler.Handle(err)
	}
}

// ApplicationErrorHandler обработчик остальных ошибок приложения
type ApplicationErrorHandler struct {
	BaseHandler
}

func (h *ApplicationErrorHandler) Handle(err error) {
	fmt.Println("Application error:", err)
}
