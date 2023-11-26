package todolist

import (
	todo "github.com/tothbence9922/gotodo/internal/todo/model"
)

type TodoList struct {
	Title string
	Todos []todo.Todo
}
