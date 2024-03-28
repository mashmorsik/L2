package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет
введена команда выхода (например \quit).

*/

func main() {
	Utility()
}

func Utility() {
	fmt.Println("Welcome to shell")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(">> ")

		// Считываем введенную пользователем команду.
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println("scan error: ", err)
		}

		// Получаем команду
		cmd := scanner.Text()
		if cmd == "\\exit" {
			break
		}

		err = execCmd(cmd)
		if err != nil {
			_, err = fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
			continue
		}
	}
}

// execCmd - функция, выполняющая команды, введенные пользователем.
func execCmd(cmd string) error {
	args := strings.Fields(cmd)
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "cd":
		return cd(args)
	case "pwd":
		return pwd()
	case "echo":
		return echo(args)
	case "kill":
		return kill(args)
	case "ps":
		return ps()
	}

	return nil
}

// cd - функция для изменения текущего рабочего каталога.
func cd(args []string) error {
	if len(args) != 2 {
		return errors.New("not enough arguments")
	}

	return os.Chdir(args[1])
}

// pwd - функция для вывода текущего рабочего каталога.
func pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Println(dir)
	return err
}

// echo - функция для вывода введенной строки.
func echo(args []string) error {
	if len(args) < 2 {
		fmt.Println("nothing to echo")
		return nil
	}

	fmt.Println(args[1:])
	return nil
}

// kill - функция для отправки сигнала завершения процессу.
func kill(args []string) error {
	if len(args) != 2 {
		return errors.New("not enough arguments")
	}

	return exec.Command("kill", args[1]).Run()
}

// ps - функция для вывода информации о текущих процессах.
func ps() error {
	proc := exec.Command("ps")
	out, err := proc.Output()
	if err != nil {
		fmt.Println("unable to show running processes")
		return err
	}

	fmt.Println(string(out))
	return nil
}
