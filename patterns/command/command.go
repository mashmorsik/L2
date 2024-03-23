package main

import "fmt"

// Command интерфейс команды
type Command interface {
	Execute()
}

// Light структура для управления светом
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is off")
}

// LightOnCommand команда включения света
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

// LightOffCommand команда выключения света
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}
