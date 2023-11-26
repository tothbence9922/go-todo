package main

import (
	"encoding/json"
	"fmt"

	"github.com/tothbence9922/gotodo/internal/hello"
	status "github.com/tothbence9922/gotodo/internal/status/model"
	todo "github.com/tothbence9922/gotodo/internal/todo/model"
	todolist "github.com/tothbence9922/gotodo/internal/todoList/model"
)

func main() {
	todoInstance := todo.Todo{Id: uint64(0), Status: status.TODO, Title: "First todo", Description: "Description is here."}

	todoBin, err := json.Marshal(todoInstance)

	if err == nil {
		fmt.Println(string(todoBin))
	}

	todoArray := make([]todo.Todo, 0)
	todoArray = append(todoArray, todoInstance)

	todoList := todolist.TodoList{Todos: todoArray, Title: "First list"}

	todoListBin, err := json.Marshal(todoList)

	if err == nil {
		fmt.Println(string(todoListBin))
	}

	greeting, err := hello.Hello("World!")

	if err != nil {
		panic("Error while constructing the greeting string!")
	}

	fmt.Println(greeting)
}
