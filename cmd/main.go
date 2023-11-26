package main

import (
	"encoding/json"
	"fmt"

	"github.com/tothbence9922/gotodo/internal/hello"
	status "github.com/tothbence9922/gotodo/internal/status/model"
	"github.com/tothbence9922/gotodo/internal/todo"
	todoModel "github.com/tothbence9922/gotodo/internal/todo/model"
	"github.com/tothbence9922/gotodo/internal/todolist"
)

func main() {
	todoInstance, err := todo.Create(uint64(0), "First todo", "Description is here.", status.TODO)

	if err == nil {
		todoBin, err := json.Marshal(todoInstance)

		if err == nil {
			fmt.Println(string(todoBin))

			todoArray := make([]todoModel.Todo, 0)
			todoArray = append(todoArray, *todoInstance)

			todoList, err := todolist.Create("First list")

			if err == nil {
				todoList.Todos = todoArray

				todoListBin, err := json.Marshal(todoList)

				if err == nil {
					fmt.Println(string(todoListBin))
				}
			}
		}
	}

	greeting, err := hello.Hello("World!")

	if err != nil {
		panic("Error while constructing the greeting string!")
	}

	fmt.Println(greeting)
}
