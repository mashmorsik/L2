package main

import (
	"fmt"
	server2 "github.com/mashmorsik/L2/tasks/task_11/infrastructure/server"
)

func main() {
	server := server2.NewServer().StartServer()
	fmt.Println(server)
}
