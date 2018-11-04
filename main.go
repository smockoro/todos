package main

import (
	"fmt"
	"log"

	todo "github.com/smockoro/todoapp/todo"
)

func main() {
	conn := "DB connection"
	tc := todo.NewTodoController(conn)
	todo, err := tc.View(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todo)
}
