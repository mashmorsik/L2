package main

import (
	"fmt"
)

// SessionState определяет интерфейс состояния сеанса пользователя
type SessionState interface {
	Authenticate(username, password string) error
	AccessResource(resource string) error
	Logout() error
}

// AuthenticatedState реализует состояние аутентифицированного сеанса пользователя
type AuthenticatedState struct{}

func (s *AuthenticatedState) Authenticate(username, password string) error {
	return fmt.Errorf("already authenticated")
}

func (s *AuthenticatedState) AccessResource(resource string) error {
	fmt.Printf("Accessing resource '%s'\n", resource)
	return nil
}

func (s *AuthenticatedState) Logout() error {
	fmt.Println("Logging out")
	return nil
}

// UnauthenticatedState реализует состояние неаутентифицированного сеанса пользователя
type UnauthenticatedState struct{}

func (s *UnauthenticatedState) Authenticate(username, password string) error {
	fmt.Printf("Authenticating user '%s'\n", username)
	return nil
}

func (s *UnauthenticatedState) AccessResource(resource string) error {
	return fmt.Errorf("not authenticated")
}

func (s *UnauthenticatedState) Logout() error {
	return fmt.Errorf("not authenticated")
}

// SessionManager управляет состоянием сеанса пользователя
type SessionManager struct {
	state SessionState
}

func NewSessionManager() *SessionManager {
	return &SessionManager{state: &UnauthenticatedState{}}
}

func (sm *SessionManager) Authenticate(username, password string) error {
	err := sm.state.Authenticate(username, password)
	if err == nil {
		sm.state = &AuthenticatedState{}
	}
	return err
}

func (sm *SessionManager) AccessResource(resource string) error {
	return sm.state.AccessResource(resource)
}

func (sm *SessionManager) Logout() error {
	err := sm.state.Logout()
	if err == nil {
		sm.state = &UnauthenticatedState{}
	}
	return err
}
